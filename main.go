package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	csvFile, _ := os.Open("Play-1.txt")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'
	lineCount := 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//log.Println(line)
		if lineCount == 0 {
			for _, item := range line {
				log.Println(item)
			}
		} else {
			break
		}
		lineCount++
	}
	log.Println("That's all folks!")
}
