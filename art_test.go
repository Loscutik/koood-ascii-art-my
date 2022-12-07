package main

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
	STEP := 0
	f, err := os.Open("testfont.txt")
	errHandle(err)
	var aLetter ArtLetter
	scanner := bufio.NewScanner(f)
		
	for i := 0; i <= 9*STEP; i++ {
		scanner.Scan()
	}
	aLetter, err = readArtLetter(scanner)
	fmt.Println(aLetter)
	if err!=nil {
		b.Fatal(err)
	}

}
