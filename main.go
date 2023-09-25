package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <filename>")
		return
	} else {
		filename := os.Args[1]
		fmt.Println("Filename:", filename)
	}
	dev()

}

func dev() {}