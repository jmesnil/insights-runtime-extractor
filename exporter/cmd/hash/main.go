package main

import (
	"crypto/sha256"
	"exporter/pkg/utils"
	"fmt"
	"os"
)

func main() {
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
