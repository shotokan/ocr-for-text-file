package ocr

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)
const MAX_CHARACTERS_PER_LINE = 27

// Convert a text into a slice of strings
func ConvertTextToSlice(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	var txtlines []string
	for scanner.Scan() {
		textLine := scanner.Text()
		txtlines = append(txtlines, textLine)	
	}
	txtlines = append(txtlines, "")	
	return txtlines
}


// Convert 3x3 characters of the account to numbers
// when a number is not recognized it is convert to ?
func ConvertToNumbers(account []string) []string {
	var numbersFound []string

	for i := 0; i < MAX_CHARACTERS_PER_LINE; i = i + MAX_LINES_PER_ACOUNT {
		line1 := account[0]
		line2 := account[1]
		line3 := account[2]
		
		number, err := recognizeNumber(line1[i: i+3], line2[i: i+3], line3[i: i+3])
		if err != nil {
			log.Println(err)
			numbersFound = append(numbersFound, "?")
			continue
		}

		numbersFound = append(numbersFound, strconv.FormatInt(int64(number), 10))
	}
	return numbersFound
}

func recognizeNumber(line1, line2, line3 string) (int, error) {
	for number, numberPatterns := range CLASSIFIED_NUMBERS {
		if numberPatterns[0] == line1 && numberPatterns[1] == line2 && numberPatterns[2] == line3 {
			return number, nil
		}
	}
	return 0, fmt.Errorf("number not found")
}
