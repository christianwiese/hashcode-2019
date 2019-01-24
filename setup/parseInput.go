package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInt(v string) int {
	x, err := strconv.ParseInt(strings.Trim(v, " "), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(x)
}

type Input struct {
}

func (i *Input) Print() {
	fmt.Println("======================INPUT=======================")
	fmt.Println("==================================================")

}

func ParseInput(path string) Input {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	// inits := strings.Split(scanner.Text(), " ")

	// largestIntersection := ParseIntersection(inits[0], inits[1])
	// numVehicles := ParseInt(inits[2])
	// numRides := ParseInt(inits[3])
	// bonus := ParseInt(inits[4])
	// numSteps := ParseInt(inits[5])

	// var rides []*Ride
	// var idx int
	// for scanner.Scan() {
	// 	rides = append(rides, ParseRide(idx, scanner.Text()))
	// 	idx++
	// }

	// if len(rides) != numRides {
	// 	panic("???? wrong number of rides")
	// }

	// return Input{
	// 	Rides:       rides,
	// 	NumVehicles: numVehicles,

	// 	LargestIntersection: largestIntersection,
	// 	NumSteps:            numSteps,
	// 	PerRideBonus:        bonus,
	// }

	return Input{}
}
