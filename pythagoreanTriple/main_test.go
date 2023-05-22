package main

import (
	"testing"
)

func TestBabilonian(t *testing.T) {
	x := 13
	y := 5
	want := [3]int{144, 130, 194}
	result, err := babilonian(x, y)
	if result != want || err != nil {
		t.Fatalf(`babilonian(%q, %q) = %q %v, want match for %#q`, x, y, result, err, want)
	}
}

func TestNegBabilonian(t *testing.T) {
	x := 13
	y := 15
	want := [3]int{0, 0, 0}
	result, err := babilonian(x, y)
	if result != want {
		t.Fatalf(`babilonian(%q, %q) = %q %v, want match for %#q`, x, y, result, err, want)
	}
}
