package main

import "testing"

func TestIsPrime(t *testing.T) {
	x := 2
	want := true
	result := isPrime(x)
	if result != want {
		t.Fatalf(`isPrime(%v) = %t, want match for %t`, x, result, want)
	}
}

func TestIsNotPrime(t *testing.T) {
	x := 40
	want := false
	result := isPrime(x)
	if result != want {
		t.Fatalf(`isPrime(%q) = %t, want match for %t`, x, result, want)
	}
}
