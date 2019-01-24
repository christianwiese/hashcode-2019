package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Output []Command

type Command struct {
	droneID     int
	operation   string
	warhouseID  int
	orderID     int
	productID   int
	numProducts int
}

func Dump(out Output, file string) {
	output := fmt.Sprintf("./%s.out", strings.TrimSuffix(file, ".in"))
	f, _ := os.Create(output)
	defer f.Close()

	w := bufio.NewWriter(f)

	num := len(out)
	w.WriteString(fmt.Sprintf("%d", num))
	w.WriteString("\n")
	for i, c := range out {
		if c.operation == "D" {
			w.WriteString(fmt.Sprintf("%d %s %d %d %d", c.droneID, c.operation, c.orderID, c.productID, c.numProducts))
		} else {
			w.WriteString(fmt.Sprintf("%d %s %d %d %d", c.droneID, c.operation, c.warhouseID, c.productID, c.numProducts))
		}

		if i < num {
			w.WriteString("\n")
		}
	}
	w.Flush()
}
