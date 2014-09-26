package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("\nPrints the Smallest Multiple.\nUsage: go run main.go NUM\n")
		return
	}

	param, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("\nPlease input numeric value only.\n")
		return
	}

	prev := 1
	for x := 1; x <= param; x++ {
		prev = lcm(x, prev)
	}

	fmt.Printf("Smallest Multiple from 1 to %[1]d is %[2]d\n", param, prev)
}
