package ocr

import (
	"bufio"
	"io"
)

// Convert it into a slice of strings
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

