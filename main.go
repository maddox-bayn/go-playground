package main

import (
	"fmt"
	"os"

	"github.com/maddox-bayn/go-playground/strutils"
)

func main() {
	// get input from cli
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <word>")
		return
	}

	word := os.Args[1]
	fmt.Printf("Word: %v\n", word)
	fmt.Printf("RuneCount: %d\n", strutils.RuneCount(word))
	fmt.Printf("FirstRune: %q\n", strutils.FirstRune(word))
	fmt.Printf("LastRune: %q\n", strutils.LastRune(word))
}
