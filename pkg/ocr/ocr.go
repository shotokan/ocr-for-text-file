package ocr

import (
	"fmt"
	"io"
	"strings"
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
	  msg := CheckSum(account)

		fmt.Println(strings.Join(account, ""), msg)
	}
}