package main

import (
	"fmt"
	"os"
)

func main() {
	input := ParseInput(os.Args[1])

	fmt.Println("WORKING ON: ", os.Args[1])

	input.Print()

	out := []Command{
		Command{droneID: 1, operation: "D"},
		Command{droneID: 2, operation: "L"},
	}
	Dump(out, os.Args[1])
}
