package main

import (
	"fmt"
	"os"
)

func main() {
	if err := skunk(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
