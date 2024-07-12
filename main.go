package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.2.0"

func main() {
	versionFlag := flag.Bool("version", false, "Display the version number")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("skunk version %s\n", version)
		os.Exit(0)
	}

	if err := skunk(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
