package domain

import (
	"log"
	"strconv"
)

const (
	OK_ACCOUNT = "OK"
	INVALID_ACCOUNT = "ERR"
	ILLEGIBLE_ACCOUNT = "ILL"
)

// Validate the account
// when 0 it is a valid account
// when different to 0 is not a valid account
// when a ? character is found the account is illigal
func checkSum(accountNumbers string) string{
	log.Println(accountNumbers)
	checkSum := 0
	delta := 9

	if len(accountNumbers) != 9 {
		errorMsg := "it is not possible to convert " + accountNumbers
		log.Println(errorMsg)
		return ILLEGIBLE_ACCOUNT
	}
	
	for i := 0; i < 9; i++ {
		accountNumber, err := strconv.Atoi(string(accountNumbers[i]))
		if err != nil {
			errorMsg := "it is not possible to convert " + string(accountNumbers[i])
			log.Println(errorMsg)
			return ILLEGIBLE_ACCOUNT
		}
		checkSum += accountNumber * delta
		delta--
	}

	checkSum = checkSum % 11
	if checkSum != 0 {
		return INVALID_ACCOUNT
	}
	
	return OK_ACCOUNT
}