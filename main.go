package main

import (
	"log"

	"./playlist"
)

func main() {
	fileName := "All-title.txt" // File format is UTF 16 LE
	maxLines := -1              // -1 are all items, > 0 is a value that cuts the export file
	pl := playlist.PlaylistRnd{}

	pl.ReadFile(fileName)
	//arr := []int{36, 37} // Export only this indexes
	//pl.SetFinalIx(arr)

	//pl.SelectItemsWithComment() // export only songs with comments

	pl.ShuffleFinalIx()

	pl.WriteFile("Randomized.txt", maxLines)

	log.Println("That's all folks!")
}
