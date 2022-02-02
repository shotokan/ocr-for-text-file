package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shotokan/ocr-for-text-file/pkg/ocr"
)

func main(){
	file, err := os.Open("example.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	x := ocr.ReadText(file)

	fmt.Println(ocr.TrainOCR(x))
}
