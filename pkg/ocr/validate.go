package ocr

import (
	"fmt"
	"log"
	"strconv"
)

func CheckSum(accountNumbers []string) (int, error){
	log.Println(accountNumbers)
	checkSum := 0
	delta := 9
	for i := 0; i < 9; i++ {
		accountNumber, err := strconv.Atoi(accountNumbers[i])
		if err != nil {
			errorMsg := "It is not possible to convert " + accountNumbers[i]
			log.Println(errorMsg)
			return 0, fmt.Errorf(errorMsg)
		}
		checkSum += accountNumber * delta
		delta--
	}

	checkSum = checkSum % 11
	return checkSum, nil
}