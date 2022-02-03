package main

import (
	"log"

	"github.com/shotokan/ocr-for-text-file/cmd/service/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}