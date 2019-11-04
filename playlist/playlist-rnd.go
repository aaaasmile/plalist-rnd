package playlist

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
	outLines  [][]string
	finalIx   []int
	fieldDet  []map[string]string
	fieldKeys []string
}

func (pl *PlaylistRnd) ReadFile(fileName string) {
	pl.outLines = [][]string{}
	pl.finalIx = []int{0}
	pl.fieldDet = []map[string]string{}
	csvFile, err := utfutil.OpenFile(fileName, utfutil.UTF16LE)
	if err != nil {
		log.Fatalln("Error on open utf16 file", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = '\t'
	lineCount := 0
	pl.fieldKeys = []string{}
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
				pl.fieldKeys = append(pl.fieldKeys, item)
			}
		} else {
			// if lineCount > 30 {
			// 	break
			// }
			mmF := make(map[string]string)
			for i, item := range line {
				k := pl.fieldKeys[i]
				mmF[k] = item
			}
			pl.fieldDet = append(pl.fieldDet, mmF)
			pl.finalIx = append(pl.finalIx, lineCount)
		}
		pl.outLines = append(pl.outLines, line)
		lineCount++
	}
	// TODO: crea il pl.finalIx in modo random
	log.Printf("Recongnized %d songs", len(pl.outLines)-1)
	pl.printFieldName()
	pl.printAllComments()
}

func (pl *PlaylistRnd) printAllComments() {
	log.Println("Print comments")
	for i, mm := range pl.fieldDet {
		com := mm["Kommentar"] // German filed title for comment (language is set inside itunes)
		if com != "" {
			log.Printf("Comment in [%d] is %q", i+1, com)
		}
	}
}

func (pl *PlaylistRnd) printFieldName() {
	log.Println("Field Name")
	for i, item := range pl.fieldKeys {
		log.Printf("[%d] %s", i, item)
	}
}

func (pl *PlaylistRnd) RemoveComments() {

}

func (pl *PlaylistRnd) SetFinalIx(arr []int) {
	pl.finalIx = []int{0} // title field is always there
	for _, item := range arr {
		pl.finalIx = append(pl.finalIx, item)
	}
}

func (pl *PlaylistRnd) WriteFile(fileName string, maxLines int) {
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
	count := 0
	for i, lineIx := range pl.finalIx {
		//fmt.Println(lineIx)
		line := pl.outLines[lineIx]
		strLineOut = ""
		if i > maxLines && maxLines > 0 {
			log.Printf("NOTE: output file cutted to %d lines", maxLines)
			break
		}
		if i > 0 {
			strLineOut = "\r\n"
			count++
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
	log.Printf("File created: %s with %d items", fileName, count)
}
