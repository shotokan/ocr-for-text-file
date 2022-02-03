package ocr

import (
	"io"
	"strings"
)

func Recognize(inputData io.Reader) []string{
	accountsRaw := ConvertTextToSlice(inputData)
	accountByLines := ParseAccounts(accountsRaw)
	accountNumbers := make([]string, 0)

	for _, accountRaw := range accountByLines {
		accountNumber := ConvertToNumbers(accountRaw)
		accountNumbers = append(accountNumbers, strings.Join(accountNumber, ""))
	}

	return accountNumbers
}