package main

import "testing"

func TestFoo(t *testing.T) {
	if foo() != "bar" {
		t.Error("Test failed")
	}
}
