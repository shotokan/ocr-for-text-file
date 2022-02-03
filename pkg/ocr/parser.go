package ocr

import (
	"fmt"
	"log"
	"strconv"
)

var Patterns map[int][3]string

func init() {
	Patterns = make(map[int][3]string)
}

// Split and transform lines of text in accounts form by 3 lines
// Each account is formed by 9 numbers and each number number is
func ParseAccounts(rawAccounts []string) [][]string{
	var accountByLines []string = make([]string, 0, 3)
	accounts := make([][]string, 0, (len(rawAccounts)/4) + 1)

	for _, textLine := range rawAccounts {
	
		if len(accountByLines) == 3 {
			accounts = append(accounts, accountByLines)
			accountByLines = []string{}
			continue
		}
		accountByLines = append(accountByLines, textLine)
		
	}
	return accounts
}

func ParseNumbers(account []string) []string{
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
	for number, numberPatterns := range Patterns {
		if numberPatterns[0] == line1 && numberPatterns[1] == line2 && numberPatterns[2] == line3 {
			return number, nil
		}
	}
	return 0, fmt.Errorf("Number not found")
}

func TrainOCR(numbers []string) map[int][3]string {
	for i, line := range ParseAccounts(numbers){
		Patterns[i] = [3]string{line[0], line[1], line[2]}
	}
	return Patterns
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