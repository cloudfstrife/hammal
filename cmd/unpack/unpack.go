package pack

import (
	"bytes"
	"compress/zlib"
	"encoding/base32"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cloudfstrife/hammal/cmd"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RegistCmd(unpackCommand)
}

var unpackCommand = &cobra.Command{
	Use:   "unpack",
	Short: "unpack read bytes from stdin ,do base32 decode ,and uncompress it , write result to stdout",
	Long:  `unpack read bytes from stdin ,do base32 decode ,and uncompress it , write result to stdout`,
	Run:   Run,
}

//Run run command
func Run(cmd *cobra.Command, args []string) {
	var (
		err     error
		decoded []byte
	)

	inText, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Error("read from stdin", err)
	}

	if decoded, err = base32.StdEncoding.DecodeString(string(inText)); err != nil {
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

//UnCompress 解压缩
func UnCompress(text []byte) ([]byte, error) {
	var (
		err error
	)
	rbuf := bytes.NewBuffer(text)
	zlibR, err := zlib.NewReader(rbuf)
	if err != nil {
		return nil, fmt.Errorf("Create zlib Reader : %w", err)
	}
	defer zlibR.Close()

	wbuf := bytes.NewBuffer([]byte{})
	if _, err = io.Copy(wbuf, zlibR); err != io.ErrUnexpectedEOF {
		return nil, fmt.Errorf("Copy Buffer : %w", err)
	}

	result := make([]byte, wbuf.Len())
	if _, err = wbuf.Read(result); err != nil {
		return nil, fmt.Errorf("Read Buffer : %w", err)
	}

	return result, nil
}
