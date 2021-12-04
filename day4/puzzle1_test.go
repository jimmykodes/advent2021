package main

import (
	"testing"
)

func Test_p1(t *testing.T) {
	d := getData("test.txt")
	want := 4512
	if got := p1(d); got != want {
		t.Errorf("expected scores to match. got %d. want %d", got, want)
	}
}
