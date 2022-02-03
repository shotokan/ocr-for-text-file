package ocr

import (
	"io"
	"log"
	"os"
)
var CLASSIFIED_NUMBERS map[int][3]string
var TRAINING_TEXT_PATH string = "./training_text.txt"

func init() {
	CLASSIFIED_NUMBERS = make(map[int][3]string)
	openedFile := readTrainingText()
	trainingText := ConvertTextToSlice(openedFile)
	trainOCR(trainingText)
}

func trainOCR(numbers []string) {
	for i, line := range ParseAccounts(numbers){
		if len(line) != 3 {
			errMsg := "OCR could not be trained"
			log.Fatalf(errMsg)
		}

		CLASSIFIED_NUMBERS[i] = [3]string{line[0], line[1], line[2]}
	}
}

func readTrainingText() io.Reader {
	file, err := os.Open(TRAINING_TEXT_PATH)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	return file
}