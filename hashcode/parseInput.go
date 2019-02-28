package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInt(v string) int {
	x, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return int(x)
}

type Input struct {
	rows       int
	columns    int
	numDrones  int
	deadLine   int
	maxLoad    int
	warehouses []Warehouse
	orders     []Order
}

type Warehouse struct {
	x        int
	y        int
	products []int
}

type Order struct {
	x        int
	y        int
	products []int
}

func ParseInput(path string) Input {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	inits := strings.Split(scanner.Text(), " ")

	fmt.Println(inits)

	rows := ParseInt(inits[0])
	columns := ParseInt(inits[1])
	numDrones := ParseInt(inits[2])
	deadLine := ParseInt(inits[3])
	maxLoad := ParseInt(inits[4])

	scanner.Scan()
	_ = ParseInt(scanner.Text())

	scanner.Scan()

	var productWeights []int
	pr := strings.Split(scanner.Text(), " ")

	for _, s := range pr {
		productWeights = append(productWeights, ParseInt(s))
	}

	scanner.Scan()

	numWarehouses := ParseInt(scanner.Text())

	var warehouses []Warehouse
	for i := 0; i < numWarehouses; i++ {
		scanner.Scan()
		xy := strings.Split(scanner.Text(), " ")
		scanner.Scan()
		p := strings.Split(scanner.Text(), " ")
		var pr []int
		for _, s := range p {
			pr = append(pr, ParseInt(s))
		}
		w := Warehouse{
			x:        ParseInt(xy[0]),
			y:        ParseInt(xy[1]),
			products: pr,
		}
		warehouses = append(warehouses, w)
	}

	scanner.Scan()

	numOrders := ParseInt(scanner.Text())

	var orders []Order
	for i := 0; i < numOrders; i++ {
		scanner.Scan()
		xy := strings.Split(scanner.Text(), " ")

		scanner.Scan()
		_ = ParseInt(scanner.Text())

		scanner.Scan()
		var pr []int
		p := strings.Split(scanner.Text(), " ")
		for _, s := range p {
			pr = append(pr, ParseInt(s))
		}
		o := Order{
			x:        ParseInt(xy[0]),
			y:        ParseInt(xy[1]),
			products: pr,
		}
		orders = append(orders, o)
	}

	return Input{
		rows:       rows,
		columns:    columns,
		numDrones:  numDrones,
		deadLine:   deadLine,
		maxLoad:    maxLoad,
		warehouses: warehouses,
		orders:     orders,
	}
}
