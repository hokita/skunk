package main

import (
	"fmt"
)

func skunk() error {
	fmt.Println(fart())
	return nil
}

func fart() string {
	return "Boom!"
}
