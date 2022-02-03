package ocr

import (
	"fmt"
	"log"
	"strconv"
)


var MAX_LINES_PER_ACOUNT = 3

// Split and transform lines of text in accounts form by 3 lines
// Each account is formed by 9 numbers and each number is a 3x3 characters
func ParseAccounts(rawAccounts []string) [][]string{
	var accountByLines []string = make([]string, 0, MAX_LINES_PER_ACOUNT)
	numberAccounts := (len(rawAccounts)/4) + 1
	
	accounts := make([][]string, 0, numberAccounts)

	for _, textLine := range rawAccounts {
	
		if len(accountByLines) == MAX_LINES_PER_ACOUNT {
			accounts = append(accounts, accountByLines)
			accountByLines = []string{}
			continue
		}
		accountByLines = append(accountByLines, textLine)
		
	}
	return accounts
}

// Converts 3x3 characters of the account to numbers
func ConvertToNumbers(account []string) []string {
	var numbersFound []string

	for i := 0; i < 27; i = i + 3 {
		line1 := account[0]
		line2 := account[1]
		line3 := account[2]
		number, err := findNumber(line1[i: i+3], line2[i: i+3], line3[i: i+3])
		if err != nil {
			log.Println(err)
			numbersFound = append(numbersFound, "?")
			continue
		}
		numbersFound = append(numbersFound, strconv.FormatInt(int64(number), 10))
	}
	return numbersFound
}

func findNumber(line1, line2, line3 string) (int, error) {
	for number, numberPatterns := range CLASSIFIED_NUMBERS {
		if numberPatterns[0] == line1 && numberPatterns[1] == line2 && numberPatterns[2] == line3 {
			return number, nil
		}
	}
	return 0, fmt.Errorf("Number not found")
}



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