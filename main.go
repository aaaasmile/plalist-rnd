package main

import (
	"log"

	"./playlist"
)

func main() {
	fileName := "All-title.txt" // File format is UTF 16 LE
	maxLines := -1              //2
	pl := playlist.PlaylistRnd{}

	pl.ReadFile(fileName)
	//arr := []int{36, 37} // esporta proprio questi indici
	//pl.SetFinalIx(arr)
	//pl.SelectItemsWithComment()
	pl.WriteFile("Randomized.txt", maxLines)

	log.Println("That's all folks!")
}
