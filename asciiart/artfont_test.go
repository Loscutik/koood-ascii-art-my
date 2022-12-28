package asciiart

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func errHandle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}



func BenchmarkReadArtLetter(b *testing.B) {
	STEP := 24
	f, err := os.Open("fonts/standard.txt")
	errHandle(err)
	var aLetter ArtString
	scanner := bufio.NewScanner(f)

	for i := 0; i <= 9*STEP; i++ {
		scanner.Scan()
	}
	aLetter, err = readArtChar(scanner)
	fmt.Println()
	aLetter.ArtPrint()
	if err != nil {
		b.Fatal(err)
	}
}
