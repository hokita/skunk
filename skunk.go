package main

import "fmt"

func skunk() error {
	fmt.Println(foo())
	return nil
}

func foo() string {
	return "foo"
}
