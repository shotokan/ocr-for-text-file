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
	accountNumbers := make([][]string, 0, len(accountByLines))
	for _, accountRaw := range accountByLines {
		accountNumber := ocr.ParseNumbers(accountRaw)
		accountNumbers = append(accountNumbers, accountNumber)
	}

	for _, account := range accountNumbers {
		checkSum, err := ocr.CheckSum(account)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(checkSum)
	}
}
