package main

import (
	"log"

	"./playlist"
)

func main() {
	fileName := "All-title.txt" // File format is UTF 16 LE
	maxLines := 2
	pl := playlist.PlaylistRnd{}
	pl.ReadFile(fileName)
	pl.WriteFile("Randomized.txt", maxLines)

	log.Println("That's all folks!")
}
