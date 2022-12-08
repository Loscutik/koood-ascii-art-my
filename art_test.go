package main

import (
	"testing"
)

func TestIsAsciiString(t *testing.T) {
	str := "fggh"
	res, runes := IsAsciiString(str)
	if !res {
		t.Fatalf("wrong runes %c", runes)
	}
}
