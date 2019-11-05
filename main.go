package main

import (
	"flag"
	"log"

	"./playlist"
)

func main() {
	var inputFile = flag.String("in", "All-title.txt", "Playlist with all songs as input (itunes export)")
	var outFile = flag.String("out", "Rnd-Grigio.txt", "Shuffled playlist output (itunes import)")

	flag.Parse()

	//fileName := "All-title.txt" // File format is UTF 16 LE
	maxLines := -1 // -1 are all items, > 0 is a value that cuts the export file
	pl := playlist.PlaylistRnd{}

	pl.ReadFile(*inputFile)
	//arr := []int{36, 37} // Export only this indexes
	//pl.SetFinalIx(arr)

	//pl.SelectItemsWithComment() // export only songs with comments

	pl.ShuffleFinalIx()

	pl.WriteFile(*outFile, maxLines)

	log.Println("That's all folks!")
}
