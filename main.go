package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var errInvalidArgument = errors.New("Invalid argument")

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		printHelp()
		os.Exit(1)
	}

	x, y, err := parseArguments(args)

	if err != nil {
		printHelp()
		os.Exit(2)
	}

	fmt.Printf("The GCD for %d and %d is %d\n", x, y, gcd(x, y))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func printHelp() {
	fmt.Println("Please give two natural numbers")
	fmt.Println("Example: gcdgo 15 9")
}

func parseArguments(args []string) (int, int, error) {
	parseArgument := func(arg string) (int, error) {
		parsedArg, err := strconv.ParseInt(arg, 10, 64)

		if err != nil {
			return 0, errInvalidArgument
		}

		if parsedArg < 0 {
			return 0, errInvalidArgument
		}

		return int(parsedArg), nil
	}

	x, err := parseArgument(args[0])

	if err != nil {
		return 0, 0, errInvalidArgument
	}

	y, err := parseArgument(args[1])

	if err != nil {
		return 0, 0, errInvalidArgument
	}

	return x, y, nil
}
