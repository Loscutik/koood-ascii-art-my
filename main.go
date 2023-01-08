package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"01.kood.tech/git/obudarah/ascii-art-my/asciiart"
)

type input struct {
	output         string
	align          string
	color          string
	lettersToColor string
	reverse        string
	banner         string
	text           string
}

// TODO reverse
func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	if args.reverse != "" {
		aFont, err := asciiart.GetArtFont("banners/" + asciiart.F_STANDART + ".txt")
		if err != nil {
			log.Fatalf("cannot get artfont: %s", err)
		}

		aText, err := getArtText(args.reverse)
		if err != nil {
			log.Fatal(err)
		}

		str, err := convertArtText(aText, aFont)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(str)
		return
	}

	// check that argument only contains ascii symbols from ' ' to '~'
	if ok, runes := asciiart.IsAsciiString(args.text); !ok {
		log.Fatalf("the programme only works with ascii symbols. Next symbols cannot be handle %v ", runes)
	}

	var aStrs []asciiart.ArtString // for keeping lines of art Text
	var CSIcolor string

	if args.lettersToColor != "" || args.align != "left" {
		aStrs, err = getStyleArtText(&args)
		if err != nil {
			log.Fatalln(err)
		}
	} else { // no letters to color or no color at all
		if args.color != "" {
			CSIcolor = getColorCode(args.color)
			fmt.Print(CSIcolor)
		}
		aStrs, err = asciiart.GetArtText(args.text, "banners/"+args.banner+".txt")
		if err != nil {
			log.Fatalln(err)
		}
	}

	// output the result into a file or terminal
	if args.output == "" {
		asciiart.ArtPrint(aStrs)
	} else {
		file, err := os.Create(args.output)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		asciiart.ArtFprint(file, aStrs)
	}
}

/*
gets the terminal width
*/
func getTerminalWidth() (int, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	col, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("cannot get the terminal's width: %s", err)
	}

	wT, err := strconv.Atoi((strings.TrimSuffix(string(col), "\n"))) // terminal's width
	if err != nil {
		return 0, fmt.Errorf("cannot get the terminal's width: %s", err)
	}

	return wT, nil
}

/*
converts text into colored artstring using given banner.
CSI is an escape sequence for setting display attributes, letters is a string which defines letters for coloring.
letters must be the given text with addition symbols befor and after the  symbols for coloring.
In purpose to avoid any ambiguity, this addition symbols must not be the same as the near symbols.
n is a length of each art ascii string counted only by printable symbols
*/
func getStyleArtText(args *input) (aText []asciiart.ArtString, err error) {
	// get  ascii art font from banner
	aFont, err := asciiart.GetArtFont("banners/" + args.banner + ".txt")
	if err != nil {
		err = fmt.Errorf("cannot get artfont: %s", err)
		return
	}

	termWidth, err := getTerminalWidth()
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(args.text, "\\n")
	CSI := getColorCode(args.color)
	needStyle := (args.lettersToColor == "")
	runesToColor := []rune(args.lettersToColor)
	iL := 0 // counter for charsLetters in the lettersToColor

	for _, str := range strs {

		if str == "" {
			aText = append(aText, asciiart.StringToArt("", aFont))
			continue
		}

		// calculate the length of the feature art string
		lenArtString := 0
		for _, ch := range str {
			lenArtString += len(aFont[ch][0])
		}

		words := strings.Split(str, " ")
		var aStr asciiart.ArtString // to form art string from the current str

		// calculate and add offsets for the aligns "right" and "center" or number of words (no spare spaces) for justify
		offset := termWidth - lenArtString
		wordsNumber := 0
		switch args.align {
		case "right":
			addOffset(&aStr, offset)
		case "center":
			addOffset(&aStr, offset/2)
		case "justify":
			for _, w := range words {
				if w != "" {
					wordsNumber++
				}
			}
		}

		setStyle(&aStr, needStyle, CSI)

		for i, word := range words {
			// look for different characters in text and lettersToColor. If they are different, change color (set or drop it)
			if word != "" {
				iCh := 0 // counter for characters in the word
				chars := []rune(word)
				for iCh < len(chars) {
					if iL < len(runesToColor) && chars[iCh] != runesToColor[iL] {
						needStyle = !needStyle
						setStyle(&aStr, needStyle, CSI)
						iL++

					} else {
						aStr.AddChar(chars[iCh], aFont)
						iCh++
						iL++
					}
				}

				if args.align == "justify" && wordsNumber != 1 {
					// calculete and add the next offset
					wordsNumber--
					currentOffset := offset / wordsNumber
					offset -= currentOffset
					addOffset(&aStr, currentOffset)
				}
			}

			if i < len(words)-1 { // not need a space after the last word
				aStr.AddChar(' ', aFont)
			}
		}
		aText = append(aText, aStr)
	}

	return
}

/*
adds CSI to aStr if it needs or drop all styles off
*/
func setStyle(aStr *asciiart.ArtString, need bool, CSI string) {
	if need {
		aStr.AddConstString(CSI)
	} else {
		aStr.AddConstString("\033[0m") // all styles off
	}
}

/*
adds add offset to ascii string
*/
func addOffset(aStr *asciiart.ArtString, offset int) {
	if offset > 0 {
		aStr.AddConstString("\033[" + strconv.Itoa(offset) + "C")
	}
}

/*
gets graphic representation from a file
*/
func getArtText(fileName string) ([][][]byte, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	splittedBytes := bytes.Split(b, []byte{'\n'})
	if splittedBytes == nil {
		return nil, nil
	}

	var res [][][]byte // 1[] - strings of artText, 2[] - lines of art string, 3[] - bytes in one line
	aString := make([][]byte, 0, asciiart.SYMBOL_HEIGHT)
	lineCounter := 0
	l := 0
	for _, s := range splittedBytes {
		if lineCounter == 0 {
			l = len(s)
		}

		if l == 0 {
			res = append(res, [][]byte{})
			continue
		}

		if len(s) != l {
			return res, fmt.Errorf("lines has different length")
		}

		aString = append(aString, s)
		lineCounter++

		if lineCounter == asciiart.SYMBOL_HEIGHT {
			res = append(res, aString)
			aString = make([][]byte, 0, asciiart.SYMBOL_HEIGHT)
			lineCounter = 0
		}
	}

	if lineCounter != asciiart.SYMBOL_HEIGHT && lineCounter != 0 {
		return res, fmt.Errorf("wrong length of art string: %d, need %d, ", lineCounter, asciiart.SYMBOL_HEIGHT)
	}

	return res, nil
}

/*
converts the graphic representation into a text.
*/
func convertArtText(aText [][][]byte, aFont asciiart.ArtFont) (string, error) {
	text := ""
	for _, aString := range aText {
		str, err := convertArtString(aString, aFont)
		if err != nil {
			return text, err
		}
		text += str + "\n"
	}
	return text, nil
}

/*
converts the graphic represented string into a normal string.
*/
func convertArtString(aString [][]byte, aFont asciiart.ArtFont) (string, error) {
	if len(aString) == 0 {
		return "", nil
	}

	EmptyArtString := asciiart.ArtString{"", "", "", "", "", "", "", ""}
	aLetter := EmptyArtString
	str := "" // converted string
	for j := 0; j < len(aString[0]); j++ {
		spaces := 0
		for i := 0; i < asciiart.SYMBOL_HEIGHT; i++ {
			aLetter[i] += string(aString[i][j])
			if aString[i][j] == ' ' {
				spaces++
			}
		}

		if spaces == asciiart.SYMBOL_HEIGHT {
			l, err := convertArtChar(aLetter, aFont, &j)
			if err != nil {
				return str, err
			}
			str += string(l)

			aLetter = EmptyArtString
		}

	}

	return str, nil
}

/*
converts the graphic represented character into a normal character.
*/
func convertArtChar(aLetter asciiart.ArtString, aFont asciiart.ArtFont, j *int) (rune, error) {
	const SPICE_LENGTH = 5

	if len(aLetter[0]) == 1 { // a column of spaces
		*j += SPICE_LENGTH // the next symbols must be spaces
		return ' ', nil
	}

	letter, err := searchLetter(aLetter, aFont)
	if err != nil {
		return letter, fmt.Errorf("incorrect graphic representation: %s", err)
	}
	return letter, nil
}

/*
returns a symbol represented by the graphic character
*/
func searchLetter(aLetter asciiart.ArtString, aFont asciiart.ArtFont) (rune, error) {
	for ch := asciiart.FIRST_SYMBOL; ch <= asciiart.LAST_SYMBOL; ch++ {
		if asciiart.IsEqual(aLetter, aFont[ch]) {
			return ch, nil
		}
	}

	// create an error
	var b strings.Builder
	aLetter.ArtFprint(&b) // form 1 string from all lines of artLetter
	return 0, fmt.Errorf("unknown symbol %s", b.String())
}
