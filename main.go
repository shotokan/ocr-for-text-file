package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shotokan/ocr-for-text-file/pkg/ocr"
)

func main(){
	file, err := os.Open("example.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	inputData, err := os.Open("entradas.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	corpus := ocr.ReadText(file)
	accountsRaw := ocr.ReadText(inputData)
	ocr.TrainOCR(corpus)
	accountByLines := ocr.ParseAccounts(accountsRaw)
	for _, accountRaw := range accountByLines {
		fmt.Println(ocr.ParseNumbers(accountRaw))

	}
}
