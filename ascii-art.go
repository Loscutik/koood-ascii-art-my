package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"01.kood.tech/git/obudarah/ascii-art-my/fonts"
)

const (
	ARGS       = 1
	F_STANDART = "standard.txt"
)

type (
	ArtFont   = map[rune]ArtString
	ArtString = [fonts.SYMBOL_HEIGHT]string
)

func main() {
	if len(os.Args) != ARGS+1 {
		log.Fatalf("the programms needs the strict %d argument", ARGS)
	}

	// check that argument only contains ascii symbols from ' ' to '~'
	if ok, runes := IsAsciiString(os.Args[1]); !ok {
		log.Fatalf("the programme only works with ascii symbols. Next symbols cannot be handle %v ", runes)
	}

	// get the beautiful font for art-printing
	aFont, err := fonts.GetArtFont("fonts/" + F_STANDART)
	if err != nil {
		log.Fatalln(err)
	}

	strs := strings.Split(os.Args[1], "\\n")
	for _, str := range strs {
		artStr, err := StringToArt(str, aFont)
		if err != nil {
			log.Fatal(err)
		}

		ArtPrint(artStr)
	}
}

/*
turns a string into a ascii graphic string
*/
func StringToArt(str string, afont ArtFont) (aStr ArtString, err error) {
	// Art string contains 8 lines. Add lines from the all string's characters: the first line of all characters, then second, and so on
	for i := 0; i < fonts.SYMBOL_HEIGHT; i++ {
		for _, ch := range str {
			aStr[i] += afont[ch][i] // Add into the i-th line of the Art String i-th line of the next character
		}
	}
	return
}

/*
checks if string contains only printable ascii symbols
*/
func IsAsciiString(str string) (bool, []rune) {
	res := true
	var notValidRunes []rune
	for _, rune := range str {
		if rune < fonts.FIRST_SYMBOL || rune > fonts.LAST_SYMBOL {
			res = false
			notValidRunes = append(notValidRunes, rune)
		}
	}
	return res, notValidRunes
}

/*
prints a ascii graphic string
*/
func ArtPrint(aStr ArtString) {
	// the empty string must comprise only 1 line
	if aStr[0] == "" {
		fmt.Println()
		return
	}

	for _, line := range aStr {
		fmt.Println(line)
	}
}
