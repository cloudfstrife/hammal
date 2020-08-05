package pack

import (
	"bytes"
	"compress/zlib"
	"encoding/base32"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cloudfstrife/hammal/cmd"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RegistCmd(packCommand)
}

var packCommand = &cobra.Command{
	Use:   "pack",
	Short: "pack read bytes from stdin ,compress it , and do base32 encode , write result to stdout",
	Long:  `pack read bytes from stdin ,compress it , and do base32 encode , write result to stdout`,
	Run:   Run,
}

//Run run command
func Run(cmd *cobra.Command, args []string) {
	var err error

	inText, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Error("read from stdin", err)
	}

	bs, err := Compress(inText)
	if err != nil {
		log.Fatal(err)
	}

	v := base32.StdEncoding.EncodeToString(bs)
	if _, err = os.Stdout.Write([]byte(v)); err != nil {
		log.Fatal("write to stdout", err)
	}
}

//Compress 压缩
func Compress(text []byte) ([]byte, error) {
	var (
		err   error
		blist []byte
	)
	wBuf := bytes.NewBuffer(blist)
	zlibW := zlib.NewWriter(wBuf)
	defer zlibW.Close()

	if _, err = zlibW.Write(text); err != nil {
		return nil, fmt.Errorf("Write bytes : %w", err)
	}

	if err = zlibW.Flush(); err != nil {
		return nil, fmt.Errorf("Flush zlib buffer : %w", err)
	}
	result := make([]byte, wBuf.Len())
	if _, err = wBuf.Read(result); err != nil {
		return nil, fmt.Errorf("Read zlib buffer : %w", err)
	}
	return result, err
}
