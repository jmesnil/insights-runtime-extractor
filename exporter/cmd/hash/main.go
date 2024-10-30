package main

import (
	"crypto/sha256"
	"exporter/pkg/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("No arguments\n")
		fmt.Printf("Usage: hash <list of strings to hash>\n")
		os.Exit(1)
	}

	strs := os.Args[1:]

	h := sha256.New()

	for i, str := range strs {
		hashStr := utils.HashString(true, h, str)
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%s", hashStr)
	}
	fmt.Printf("\n")

}
