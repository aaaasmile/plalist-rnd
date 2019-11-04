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
	arr := []int{37, 38} // esporta proprio questi indici
	pl.SetFinalIx(arr)
	pl.WriteFile("Randomized.txt", maxLines)

	log.Println("That's all folks!")
}
