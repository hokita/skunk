package main

import "testing"

func TestFoo(t *testing.T) {
	if fart() != "Boom!" {
		t.Error("Test failed")
	}
}
