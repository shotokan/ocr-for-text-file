package ocr

import (
	"io"
	"strings"
)

func Recognize(inputData io.Reader) map[string]string{
	accountsRaw := ConvertTextToSlice(inputData)
	accountByLines := ParseAccounts(accountsRaw)
	accountNumbers := make([][]string, 0, len(accountByLines))

	for _, accountRaw := range accountByLines {
		accountNumber := ConvertToNumbers(accountRaw)
		accountNumbers = append(accountNumbers, accountNumber)
	}

	accountsValidated := make(map[string]string)

	for _, account := range accountNumbers {
	  msg := CheckSum(account)
		accountsValidated[strings.Join(account, "")] = msg
	}
	return accountsValidated
}