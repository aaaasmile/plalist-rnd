package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"unicode/utf16"

	"github.com/TomOnTime/utfutil"
)

type PlaylistRnd struct {
	outLines [][]string
	finalIx  []int
}

func (pl *PlaylistRnd) ReadFile(fileName string) {
	pl.outLines = [][]string{}
	pl.finalIx = []int{0}
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
			// if lineCount > 30 {
			// 	break
			// }
			pl.finalIx = append(pl.finalIx, lineCount)
		}
		pl.outLines = append(pl.outLines, line)
		lineCount++
	}
	// TODO: crea il pl.finalIx in modo random
	log.Printf("Recongnized %d songs", len(pl.outLines)-1)
}

func (pl *PlaylistRnd) WriteFile(fileName string) {
	if len(pl.outLines) == 0 {
		log.Fatalln("Original playlist is empty")
	}
	var bytes [2]byte
	const BOM = '\ufffe' //LE. for BE '\ufeff'

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Can't open file. %v", err)
	}
	defer file.Close()

	bytes[0] = BOM >> 8
	bytes[1] = BOM & 255

	file.Write(bytes[0:])

	strLineOut := ""
	for i, lineIx := range pl.finalIx {
		//fmt.Println(lineIx)
		line := pl.outLines[lineIx]
		strLineOut = ""
		if i > 0 {
			strLineOut = "\r\n"
		}
		for j, data := range line {
			if j > 0 {
				strLineOut += "\t"
			}
			strLineOut += data
		}
		runes := utf16.Encode([]rune(strLineOut))
		for _, r := range runes {
			bytes[1] = byte(r >> 8)
			bytes[0] = byte(r & 255)
			file.Write(bytes[0:])
		}
	}
	log.Println("File created: ", fileName)
}

func main() {
	fileName := "All-title.txt" // File format is UTF 16 LE

	pl := PlaylistRnd{}
	pl.ReadFile(fileName)
	pl.WriteFile("Randomized.txt")

	log.Println("That's all folks!")
}
