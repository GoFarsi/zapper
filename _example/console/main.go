package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(false, zapper.WithTimeFormat(zapper.RFC3339NANO))
	if err := z.NewCore(zapper.ConsoleWriterCore(true)); err != nil {
		log.Fatal(err)
	}

	z.Info("test info")
	z.Debug("debug level")
}
