package main

import (
	_ "github.com/cloudfstrife/hammal/cmd/pack"
	_ "github.com/cloudfstrife/hammal/cmd/unpack"

	"github.com/cloudfstrife/hammal/cmd"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func main() {
	cmd.Execute()
}
