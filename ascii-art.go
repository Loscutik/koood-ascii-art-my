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
)

type (
	ArtChar   = [SYMBOL_HEIGHT]string
	ArtFont   = map[rune]ArtChar
	ArtString = [SYMBOL_HEIGHT]string
)

// TODO add handles of errors
// TODO  handle \!
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
	// for char := (' '); char <= '~'; char++ {
	// 	for i := 0; i < SYMBOL_HEIGHT; i++ {
	// 		fmt.Print(char)
	// 		fmt.Print(": ")
	// 		fmt.Println((aFont[char][i]))
	// 	}
	// }

	strs := strings.Split(os.Args[1], "\\n")
	for _, str := range strs {
		artStr, err := stringToArt(str, aFont)
		if err != nil {
			log.Fatal(err)
		}

		artPrint(artStr)
	}
}

/*
reads a file with an ascii graphic fonts representing characters from ' '(space) to '~' and returns a map which keeps the characters
maps's key is the character (type rune), map's element is a value of the type ArtChar= [SYMBOL_HEIGHT]string
*/
func GetArtFont(fileName string) (aFont ArtFont, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	aFont = make(map[rune]ArtChar)
	scanner := bufio.NewScanner(file)

	// reading all characters to the map aFont
	for char := (' '); char <= '~'; char++ {
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
func readArtChar(scanner *bufio.Scanner) (aChar ArtChar, err error) {
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
func stringToArt(str string, afont ArtFont) (aStr ArtString, err error) {
	// TODO check: rune has to be from ' ' to '~'
	// the empty string must comprise only 1 line
	if str =="" { 
		aStr[0]="\n"
		return 
	}

	// Art string contains 8 lines. Add lines from the all string's characters: the first line of all characters, then second, and so on
	for i := 0; i < SYMBOL_HEIGHT; i++ {
		for _, ch := range str {
			aStr[i] += afont[ch][i] // Add into the i-th line of the Art String i-th line of the next character
		}
		aStr[i] += "\n"
	}
	return
}

/*
prints a ascii graphic string
*/
func artPrint(aStr ArtString) {
	for _, line := range aStr {
		fmt.Print(line)
	}
}
