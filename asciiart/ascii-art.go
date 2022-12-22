package asciiart

import (
	"fmt"
	"io"
	"strings"

	"01.kood.tech/git/obudarah/ascii-art-my/fonts"
)

const (
	ARGS       = 1
	F_STANDART = "standard.txt"
)

type ArtString [fonts.SYMBOL_HEIGHT]string

/*
turns a string into an ascii graphic string
*/
func StringToArt(str string, afont fonts.ArtFont) (aStr ArtString) {
	// Art string contains 8 lines. Add lines from the all string's characters: the first line of all characters, then second, and so on
	for _, ch := range str {
		for i := 0; i < fonts.SYMBOL_HEIGHT; i++ {
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
prints an ascii graphic string
*/
func (aStr *ArtString) ArtPrint() {
	// the empty string must comprise only 1 line
	if aStr[0] == "" {
		fmt.Println()
		return
	}

	for _, line := range *aStr {
		fmt.Printf("%s\n", line)
	}
}

/*
turns a whole text into one string, which represents ascii art view of the given text. banner is a path to the file with ascii-art banner
*/
func TextToArt(text string, banner string) (aText string, err error) {
	artFont, err := fonts.GetArtFont(banner)
	if err != nil {
		err = fmt.Errorf("cannot get artfont: %s", err)
		return "", err
	}

	strs := strings.Split(text, "\\n")
	for _, str := range strs {
		if str == "" {
			aText += "\n"
		} else {
			aStr := StringToArt(str, artFont)
			for i := 0; i < fonts.SYMBOL_HEIGHT; i++ {
				aText += aStr[i] + "\n"
			}
		}
	}
	return aText, nil
}

/*
prints an ascii graphic string
*/
func (aStr *ArtString) ArtFprint(w io.Writer) {
	// the empty string must comprise only 1 line
	if aStr[0] == "" {
		fmt.Fprintln(w)
		return
	}

	for _, line := range *aStr {
		fmt.Fprint(w, line)
	}
}
