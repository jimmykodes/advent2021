package main

import (
	"testing"
)

func Test_p2(t *testing.T) {
	d := getData("test.txt")
	want := 1924
	if got := p2(d); got != want {
		t.Errorf("expected scores to match. got %d. want %d", got, want)
	}
}
