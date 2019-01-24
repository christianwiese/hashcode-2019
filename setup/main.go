package main

import (
	"fmt"
	"os"
)

func main() {
	input := ParseInput(os.Args[1])

	fmt.Println("WORKING ON: ", os.Args[1])

	input.Print()
}
