/*
A strange behaviour. When I use map which keep an array of **string** the program works properly. (the map stores an art-symbol)
But when I use **[]byte** insted of string an error occurs.
If I print a symbol just right after keeping it in map  (befor keeping the next), it's printed properly.
If I print one of first symbols from ' ' to 'E' it is a mess.
If I only store in the map symbols from ' ' to 'W' - it works properly.
If I only store in the map symbols from ' ' to ” or more - the same mess.
If I delete symbols from 'R' to 'W' from  banner file (standart.txt), it will work properly.
The problem only occur with standard.txt and shadow.txt. No problem with thinkertoy.txt
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	F_STANDART    = "standard.txt"
	SYMBOL_HEIGHT = 8
	FIRST         = ' '
	LAST          = 'X'
)

type (
	ArtLetter = [SYMBOL_HEIGHT][]byte // when I use []byte - map keeps wrong bytes, when I use string - ok
	ArtFont   = map[rune]ArtLetter
	ArtString = [SYMBOL_HEIGHT]string
)

func main() {
	aFont, _ := GetArtFont(F_STANDART)

	for char := (FIRST); char <= LAST; char++ {
		for i := 0; i < SYMBOL_HEIGHT; i++ {
			fmt.Print(char)
			fmt.Print(": ")
			fmt.Println(string(aFont[char][i]))
		}
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
	defer file.Close()
	aFont = make(map[rune]ArtLetter)
	scanner := bufio.NewScanner(file)

	for char := (FIRST); char <= LAST; char++ {
		scanner.Scan()
		aFont[char], err = readArtLetter(scanner)
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
func readArtLetter(scanner *bufio.Scanner) (aLetter ArtLetter, err error) {
	for line := 0; line < SYMBOL_HEIGHT; line++ {
		if scanner.Scan() {
			aLetter[line] = (scanner.Bytes())
		} else if err = scanner.Err(); err != nil {
			break
		} else {
			err = fmt.Errorf("unexpected break of scanning")
			break
		}
	}
	return aLetter, err
}
