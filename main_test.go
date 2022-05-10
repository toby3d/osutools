package main

import (
	"testing"
)

var (
	path = []string{
		"T:\\Test\\Songs\\test.jpg",
		"T:\\Test\\Songs\\test.mp3",
		"T:\\Test\\Songs\\test.ini",
		"T:\\Test\\Songs\\test.osu",
		"T:\\Test\\Songs\\test.htm",
		"T:\\Test\\Songs\\test.osu",
	}
	emptypath = []string{
		"",
	}
)

func TestFilePosition(t *testing.T) {
	got := filePosition(path)
	want := 3
	if got != want {
		t.Errorf("got %d. want %d", got, want)
		t.Fail()
	}
	t.Logf("got %d. want %d", got, want)
}
