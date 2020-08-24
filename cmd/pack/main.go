package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base32"
	"fmt"
	"io"
	"os"

	"log"
)

const (
	// LENGTH length
	LENGTH = 1024
)

func main() {
	var err error

	reader := bufio.NewReader(os.Stdin)
	for {
		blist := make([]byte, LENGTH)
		var n int
		if n, err = reader.Read(blist); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println("read from stdin", err)
				os.Exit(1)
			}
		}
		bs, err := Compress(blist[0:n])
		if err != nil {
			log.Fatal(err)
		}

		v := base32.StdEncoding.EncodeToString(bs)
		if _, err = os.Stdout.Write([]byte(v + "\n")); err != nil {
			log.Fatal("write to stdout", err)
		}
	}
}

//Compress compress bytes
func Compress(text []byte) ([]byte, error) {
	var (
		err   error
		blist []byte
	)
	wBuf := bytes.NewBuffer(blist)
	var glibW *gzip.Writer
	if glibW, err = gzip.NewWriterLevel(wBuf, gzip.BestCompression); err != nil {
		return nil, fmt.Errorf("Create zlib Writer : %w", err)
	}

	defer glibW.Close()

	if _, err = glibW.Write(text); err != nil {
		return nil, fmt.Errorf("Write bytes : %w", err)
	}

	if err = glibW.Flush(); err != nil {
		return nil, fmt.Errorf("Flush zlib buffer : %w", err)
	}
	result := make([]byte, wBuf.Len())
	if _, err = wBuf.Read(result); err != nil {
		return nil, fmt.Errorf("Read zlib buffer : %w", err)
	}
	return result, err
}
