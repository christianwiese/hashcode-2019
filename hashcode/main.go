package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	input := ParseInput(os.Args[1])

	fmt.Println("WORKING ON: ", os.Args[1])

	for _, order := range input.orders {

		fmt.Printf("order %v\n", order)
		for _, warhouse := range input.warehouses {
			fmt.Printf("warhouse %v\n", warhouse)
			fmt.Println(distance(order, warhouse))

			var contains bool
			for _, opr := range order.products {
				for _, wpr := range warhouse.products {
					if opr == wpr {
						contains = true
					}
				}
			}
			fmt.Println(contains)
		}
	}

	out := []Command{
		Command{droneID: 1, operation: "D"},
		Command{droneID: 2, operation: "L"},
	}
	Dump(out, os.Args[1])
}

func distance(o Order, w Warehouse) int {

	return int(math.Ceil(math.Sqrt(math.Pow(float64(o.x-w.x), 2) + math.Pow(float64(o.y-w.y), 2))))
}
