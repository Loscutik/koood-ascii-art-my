package main

import (
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
	banner         string
	text           string
}

// TODO reverse
func main() {
	args, err := parseArgs()
	if err != nil {
		log.Fatal(err)
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
	charsLettters := []rune(args.lettersToColor)
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
		//- offsets := make([]int, len(words)-1)
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
					if iL < len(charsLettters) && chars[iCh] != charsLettters[iL] {
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
