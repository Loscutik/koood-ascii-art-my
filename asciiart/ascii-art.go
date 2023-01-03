package asciiart

import (
	"fmt"
	"io"
	"strings"
)

const (
	F_STANDART = "standard"
)

/*
checks if string contains only printable ascii symbols
*/
func IsAsciiString(str string) (bool, []rune) {
	res := true
	var notValidRunes []rune
	for _, rune := range str {
		if rune < FIRST_SYMBOL || rune > LAST_SYMBOL {
			res = false
			notValidRunes = append(notValidRunes, rune)
		}
	}
	return res, notValidRunes
}

/*
turns a string into an ascii graphic string. StringToArt does not split the given string.
*/
func StringToArt(str string, afont ArtFont) (aStr ArtString) {
	// Art string contains 8 lines. Add lines from the all string's characters: the first line of all characters, then second, and so on
	for _, ch := range str {
		for i := 0; i < SYMBOL_HEIGHT; i++ {
			aStr[i] += afont[ch][i] // Add into the i-th line of the Art String i-th line of the next character
		}
	}
	return
}

/*
turns a whole text into one string, which represents ascii art view of the given text. banner is a path to the file with ascii-art banner
*/
func GetArtTextInOneString(text string, banner string) (aText string, err error) {
	aStrs, err := GetArtText(text, banner)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	ArtFprint(&b, aStrs)

	return b.String(), nil
}

/*
splits text into strings separated by "\n" and converts the slice of strings into slice of artstring
*/
func GetArtText(text string, banner string) (aText []ArtString, err error) {
	aFont, err := GetArtFont(banner)
	if err != nil {
		err = fmt.Errorf("cannot get artfont: %s", err)
		return nil, err
	}

	strs := strings.Split(text, "\\n")
	for _, str := range strs {
		if str == "" {
			aText = append(aText, StringToArt("", aFont))
		} else {
			aStr := StringToArt(str, aFont)
			aText = append(aText, aStr)
		}
	}
	return aText, nil
}

/*
adds a symbol to the ascii graphic string
*/
func (aLine *ArtString) AddChar(ch rune, afont ArtFont) {
	// Art string contains 8 lines. Add lines from the all string's characters: the first line of all characters, then second, and so on
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		aLine[i] += afont[ch][i] // Add into the i-th line of the Art String i-th line of the next character
	}
}

/*
adds the given string str to each line of the ascii art string
*/
func (aStr *ArtString) AddConstString(str string) {
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		aStr[i] += str
	}
}

/*
puts the given string str before each line of the ascii art string
*/
func (aStr *ArtString) AddPrefixConstString(str string) {
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		aStr[i] = str + aStr[i]
	}
}


/*
returns length of the ascii string, i.e. a number of bytes in any of its
*/
func (aStr *ArtString) Len() int {
	if aStr == nil {
		return 0
	}
	return len(aStr[0])
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
prints a slice of the ascii art strings
*/
func ArtPrint(aStrs []ArtString) {
	for _, aStr := range aStrs {
		aStr.ArtPrint()
	}
}

/*
prints an ascii graphic string into io.Writer
*/
func (aStr *ArtString) ArtFprint(w io.Writer) {
	// the empty string must comprise only 1 line
	if aStr[0] == "" {
		fmt.Fprintln(w)
		return
	}

	for _, line := range *aStr {
		fmt.Fprintln(w, line)
	}
}

/*
prints a slice of the ascii art strings  into io.Writer
*/
func ArtFprint(w io.Writer, aStrs []ArtString) {
	for _, aStr := range aStrs {
		aStr.ArtFprint(w)
	}
}
