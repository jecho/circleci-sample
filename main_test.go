package main

import (
	"testing"
)

func TestInputWord(t *testing.T) {
	s := "Phrase Found"
	got := InputWord(s)
	if got != "Phrase Found" {
		t.Error("Could not find: ", got)
	}

	s2 := "don't find this"
	got = InputWord(s2)
	if got == "FoundMe?" {
		t.Error("Could not find: ", got)
	}
}
