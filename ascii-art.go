package main

import (
	"bytes"
	"log"
	"os"
)

const (
	F_STANDART    = "standart.txt"
	SYMBOL_HEIGHT = 8
)

type (
	ArtFont   = map[byte][SYMBOL_HEIGHT][]byte
	ArtString = [SYMBOL_HEIGHT][]byte
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("the programms needs strict 1 argument")
	}

	// get the beautiful font for art-printing
	artFont, err := GetArtFont("fonts/" + F_STANDART)
	strs := bytes.Split([]byte(os.Args[1]), []byte("\n"))
	for _, str := range strs {
		artStr := stringToArt(str,artFont)
		artPrint(artStr)
	}
}

func GetArtFont(fileName string) (afont ArtFont, err error) {
	return
}

func stringToArt(str []byte, afont ArtFont) ArtString {
	return nil
}

func artPrint(astr ArtString) {
}
