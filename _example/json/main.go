package main

import (
	"github.com/GoFarsi/zapper"
	"log"
)

func main() {
	z := zapper.New(true, zapper.WithDebugLevel(), zapper.WithServiceDetails(23, "zapper"))
	if err := z.NewCore(zapper.JsonWriterCore("./test_data", ".json", nil)); err != nil {
		log.Fatal(err)
	}

	z.Debug("debug log")
	z.Info("info log")
	z.Warn("warn log")
	z.Error("error log")
}
