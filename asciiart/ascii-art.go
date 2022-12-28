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
turns a string into an ascii graphic string
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
adds a symbol to the ascii graphic string
*/
func (aLine *ArtString) AddChar(ch byte, afont ArtFont) {
	// Art string contains 8 lines. Add lines from the all string's characters: the first line of all characters, then second, and so on
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		aLine[i] += afont[rune(ch)][i] // Add into the i-th line of the Art String i-th line of the next character
	}
}

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
turns a whole text into one string, which represents ascii art view of the given text. banner is a path to the file with ascii-art banner
*/
func TextToArt(text string, banner string) (aText string, err error) {
	aStrs, err := GetArtText(text, banner)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	ArtFprint(&b, aStrs)

	return b.String(), nil

	// old version
	// strs := strings.Split(text, "\\n")
	//
	//	for _, str := range strs {
	//		if str == "" {
	//			aText += "\n"
	//		} else {
	//			aStr := StringToArt(str, artFont)
	//			for i := 0; i < SYMBOL_HEIGHT; i++ {
	//				aText += aStr[i] + "\n"
	//			}
	//		}
	//	}
	//
	// return aText, nil
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
			aStr.AddConstString("\n")
			aText = append(aText, aStr)
		}
	}
	return aText, nil
}

/*
converts text into colored artstring using given banner.
CSI is an escape sequence for setting display attributes, letters is a string which defines letters for coloring.
letters must be the given text with addition symbols befor and after the  symbols for coloring.
In purpose to avoid any ambiguity, this addition symbols must not be the same as the near symbols.
*/
func GetStyleArtText(text string, banner string, CSI string, letters string) (aText []ArtString, err error) {
	// get  ascii art font from banner
	aFont, err := GetArtFont(banner)
	if err != nil {
		err = fmt.Errorf("cannot get artfont: %s", err)
		return nil, err
	}

	var res []ArtString
	needStyle := false
	strs := strings.Split(text, "\\n")
	charsLettters := []byte(letters)
	iL := 0 // counter for charsLetters

	for _, str := range strs {

		if str == "" {
			res = append(res, StringToArt("", aFont))
			continue
		}

		// if set of letters for color is ended it is needed to add all next str without changing the style
		if iL > len(charsLettters) {
			res = append(res, StringToArt(str, aFont))
			continue
		}

		var aStr ArtString // art string formed from the current str
		aStr.setStyle(needStyle, CSI)
		iS := 0 // counter for str
		chars := []byte(str)
		for iS < len(chars) {
			if chars[iS] != charsLettters[iL] {
				aStr.setStyle(needStyle, CSI)
				needStyle = !needStyle
				iL++

			} else {
				aStr.AddChar(chars[iS], aFont)
				iS++
				iL++
			}
		}
		res = append(res, aStr)
	}

	return res, nil
}

/*
adds CSI to aStr if it needs or drop all styles off
*/
func (aStr *ArtString) setStyle(need bool, CSI string) {
	if need {
		aStr.AddConstString(CSI)
	} else {
		aStr.AddConstString("\033[0m") // all styles off
	}
}

func (aStr *ArtString) AddConstString(str string) {
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		aStr[i] += str
	}
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
		fmt.Fprint(w, line)
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
