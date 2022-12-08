package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ARGS          = 1
	F_STANDART    = "standard.txt"
	SYMBOL_HEIGHT = 8
	FIRST_SYMBOL  = ' '
	LAST_SYMBOL   = '~'
)

type (
	ArtFont   = map[rune]ArtString
	ArtString = [SYMBOL_HEIGHT]string
)

func main() {
	if len(os.Args) != ARGS+1 {
		log.Fatalf("the programms needs the strict %d argument", ARGS)
	}

	// get the beautiful font for art-printing
	aFont, err := GetArtFont("fonts/" + F_STANDART)
	if err != nil {
		log.Fatalln(err)
	}
	/*test print - whole font */
	//  for char := (' '); char <= '4'; char++ {
	// 		fmt.Print(char)
	// 		fmt.Println(": ")
	// 		ArtPrint(aFont[char])
	// 	}

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
reads a file with an ascii graphic fonts representing characters from ' '(space) to '~' and returns a map which keeps the characters
maps's key is the character (type rune), map's element is a value of the type ArtString= [SYMBOL_HEIGHT]string
*/
func GetArtFont(fileName string) (aFont ArtFont, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	aFont = make(map[rune]ArtString)
	scanner := bufio.NewScanner(file)

	// reading all characters to the map aFont
	for char := (FIRST_SYMBOL); char <= LAST_SYMBOL; char++ {
		scanner.Scan() // skip the first empty line

		aFont[char], err = readArtChar(scanner)
		if err != nil {
			err = fmt.Errorf("error reading the character from a file: %s; character: %c; file:  %s", err, char, fileName)
			break
		}
	}

	return
}

/*
reads the next graphic character from given scanner, will return an error if scanner.Scan finishes with an error
*/
func readArtChar(scanner *bufio.Scanner) (aChar ArtString, err error) {
	for line := 0; line < SYMBOL_HEIGHT; line++ {
		if scanner.Scan() {
			aChar[line] = (scanner.Text())
		} else if err = scanner.Err(); err != nil {
			break
		} else {
			err = fmt.Errorf("unexpected break of scanning")
			break
		}
	}
	return aChar, err
}

/*
turns a string into a ascii graphic string
*/
func StringToArt(str string, afont ArtFont) (aStr ArtString, err error) {
	if ok, runes := IsAsciiString(str); !ok {
		err = fmt.Errorf("the programme only works with ascii symbols. Next symbols cannot be handle %v ", runes)
	}
	// Art string contains 8 lines. Add lines from the all string's characters: the first line of all characters, then second, and so on
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		for _, ch := range str {
			aStr[i] += afont[ch][i] // Add into the i-th line of the Art String i-th line of the next character
		}
		// aStr[i] += "\n"
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
		if rune < FIRST_SYMBOL || rune > LAST_SYMBOL {
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
