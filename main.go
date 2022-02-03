package main

import (
	"log"
	"os"

	"github.com/shotokan/ocr-for-text-file/pkg/ocr"
)

func main(){

	inputData, err := os.Open("entradas.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	ocr.Recognize(inputData)
}
