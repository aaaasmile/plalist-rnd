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
			log.Printf("Comment in [%d] is %q", i, com)
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
	// for _, md := range pl.fieldDet {

	// }
}

func (pl *PlaylistRnd) SetFinalIx(arr []int) {
	pl.finalIx = []int{}
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

	// write BOM
	file.Write(bytes[0:])

	// write title
	strLineOut := pl.buildOutLineString(0)
	writeStringInUtf16LEFile(file, strLineOut)

	// Write data song
	count := 0
	for i, lineIx := range pl.finalIx {
		//fmt.Println(lineIx)
		if i > maxLines && maxLines > 0 {
			log.Printf("NOTE: output file cutted to %d lines", maxLines)
			break
		}
		count++
		//strLineOut = "\r\n" + pl.buildOutLineString(lineIx+1)
		strLineOut = "\r\n" + pl.buildOutLineStringFromMap(lineIx)
		writeStringInUtf16LEFile(file, strLineOut)
	}
	log.Printf("File created: %s with %d items", fileName, count)
}

func (pl *PlaylistRnd) buildOutLineStringFromMap(lineIx int) string {
	fields := pl.fieldDet[lineIx]
	strLineOut := ""
	for j, kk := range pl.fieldKeys {
		if j > 0 {
			strLineOut += "\t"
		}
		strLineOut += fields[kk]
	}
	return strLineOut
}

func (pl *PlaylistRnd) buildOutLineString(lineIx int) string {
	line := pl.outLines[lineIx]
	strLineOut := ""
	for j, data := range line {
		if j > 0 {
			strLineOut += "\t"
		}
		strLineOut += data
	}
	return strLineOut
}

func writeStringInUtf16LEFile(file *os.File, strLineOut string) {
	var bytes [2]byte
	runes := utf16.Encode([]rune(strLineOut))
	for _, r := range runes {
		bytes[1] = byte(r >> 8)
		bytes[0] = byte(r & 255)
		file.Write(bytes[0:])
	}
}
