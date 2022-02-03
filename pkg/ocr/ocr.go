package ocr

import (
	"fmt"
	"io"
	"log"
)

func Recognize(inputData io.Reader) {
	accountsRaw := ConvertTextToSlice(inputData)
	accountByLines := ParseAccounts(accountsRaw)
	accountNumbers := make([][]string, 0, len(accountByLines))

	for _, accountRaw := range accountByLines {
		accountNumber := ConvertToNumbers(accountRaw)
		accountNumbers = append(accountNumbers, accountNumber)
	}

	for _, account := range accountNumbers {
		checkSum, err := CheckSum(account)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(checkSum)
	}
}