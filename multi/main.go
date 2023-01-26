package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(false, zapper.WithDebugLevel())
	if err := z.NewCore(
		zapper.ConsoleWriterCore(true),
		zapper.FileWriterCore("./test_data", nil),
	); err != nil {
		log.Fatal(err)
	}

	z.Debug("debug log")
	z.Info("info log")
	z.Warn("warn log")
	z.Error("error log")
}
