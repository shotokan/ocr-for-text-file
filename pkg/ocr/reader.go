package ocr

import (
	"bufio"
	"io"
)

// Parse a text to a slice of strings
func ReadText(file io.Reader) []string {
	scanner := bufio.NewScanner(file)
	var txtlines []string
	for scanner.Scan() {
		textLine := scanner.Text()
		txtlines = append(txtlines, textLine)	
	}
	txtlines = append(txtlines, "")	
	return txtlines
}

