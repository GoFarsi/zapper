package main

import (
	"github.com/GoFarsi/zapper"
	"log"
	"os"
)

func main() {
	z := zapper.New(false)
	if err := z.NewCore(zapper.SentryCore(os.Getenv("DSN"), "test", nil)); err != nil {
		log.Fatal(err)
	}

	err(z)
}

func err(z zapper.Zapper) {
	z.Error("test error new")
}
