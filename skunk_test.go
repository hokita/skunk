package main

import "testing"

func TestFart(t *testing.T) {
	if fart() != "Boom!" {
		t.Error("Test failed")
	}
}
