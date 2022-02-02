package ocr

import (
	"fmt"
	"log"
)

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
		log.Println(textLine)
		accountByLines = append(accountByLines, textLine)
		
	}
	return accounts
}

func ParseNumbers(account []string) []int{
	for i := 0; i < 27; i = i + 3 {
		fmt.Println(i)
		line1 := account[0]
		line2 := account[1]
		line3 := account[2]
		fmt.Println(line1[i: i+3])
		fmt.Println(line2[i: i+3])
		fmt.Println(line3[i: i+3])
	}
	return []int{1, 2}
}

// func findNumber(patterns map[int][3]string, line1, line2, line3 string) int {
// 	for number, numberPatterns := range patterns {

// 	}
// }

func TrainOCR(numbers []string) map[int][3]string {
	numberPatterns := make(map[int][3]string)
	for i, line := range ParseAccounts(numbers){
		numberPatterns[i] = [3]string{line[0], line[1], line[2]}
	}
	return numberPatterns
}