package main

import (
	"fmt"
	"os"
	"strconv"
)

func solve(index, val, largestPF int) {

	if index > val {
		fmt.Println("Largest prime factor is ", largestPF)
		return
	}

	if val%index == 0 {
		if index > largestPF {
			largestPF = index
		}

		fmt.Println("prime: ", index)

		val = val / index

		solve(index, val, largestPF)
	} else {
		solve(index+1, val, largestPF)
	}

}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("\nPrints the largest prime factor using brute-force method.\nUsage: go run main.go NUM\n")
		return
	}

	param, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("\nPlease input numeric value only.\n")
		return
	}

	solve(2, param, 0)
}
