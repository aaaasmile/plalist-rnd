package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"

	"github.com/TomOnTime/utfutil"
)

func main() {
	fileName := "Play-1.txt" // File format is UTF 16 LE
	//csvFile, _ := os.Open("Play-1.txt")
	//data, err := ioutil.ReadFile(fileName)
	csvFile, err := utfutil.OpenFile(fileName, utfutil.UTF16LE)
	if err != nil {
		log.Fatalln("Error on open utf16 file", err)
	}
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
			for i, item := range line {
				log.Printf("[%d] %s", i, item)
			}
		} else {
			break
		}
		lineCount++
	}
	log.Println("That's all folks!")
}
