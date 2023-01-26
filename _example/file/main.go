package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(true, zapper.WithDebugLevel())
	if err := z.NewCore(zapper.FileWriterCore("./test_data", nil)); err != nil {
		log.Fatal(err)
	}

	z.Debug("debug log")
	z.Info("info log")
	z.Warn("warn log")
	z.Error("error log")
}
