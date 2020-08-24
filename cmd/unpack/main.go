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

func main() {
	var (
		err     error
		decoded []byte
	)

	reader := bufio.NewReader(os.Stdin)
	for {
		var blist []byte
		if blist, err = reader.ReadBytes('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Println("read from stdin", err)
				os.Exit(1)
			}
		}
		if decoded, err = base32.StdEncoding.DecodeString(string(blist)); err != nil {
			log.Fatalf("Base32 decode : %v", err)
		}

		var v []byte
		if v, err = UnCompress(decoded); err != nil {
			log.Fatal(err)
		}

		if _, err = os.Stdout.Write([]byte(v)); err != nil {
			log.Fatal("write to stdout", err)
		}
	}
}

//UnCompress uncompress bytes
func UnCompress(text []byte) ([]byte, error) {
	var (
		err error
	)
	rbuf := bytes.NewBuffer(text)
	gzipR, err := gzip.NewReader(rbuf)
	if err != nil {
		return nil, fmt.Errorf("Create zlib Reader : %w", err)
	}
	defer gzipR.Close()

	wbuf := bytes.NewBuffer([]byte{})
	if _, err = io.Copy(wbuf, gzipR); err != io.ErrUnexpectedEOF {
		return nil, fmt.Errorf("Copy Buffer : %w", err)
	}

	result := make([]byte, wbuf.Len())
	if _, err = wbuf.Read(result); err != nil {
		return nil, fmt.Errorf("Read Buffer : %w", err)
	}

	return result, nil
}
