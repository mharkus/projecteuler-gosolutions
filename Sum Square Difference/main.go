package main

import (
	"fmt"
	"os"
	"strconv"
)

func solveSumOfSquares(cs chan int, max, n, sum int) {
	if n > max {
		cs <- sum
		return
	}

	solveSumOfSquares(cs, max, n+1, sum+n*n)
}

func solveSquareOfSums(cs chan int, max, n, sum int) {
	if n > max {
		cs <- sum * sum
		return
	}

	solveSquareOfSums(cs, max, n+1, sum+n)
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("\nPrints the Sum Square Difference.\nUsage: go run main.go NUM\n")
		return
	}

	param, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("\nPlease input numeric value only.\n")
		return
	}

	pipe := make(chan int)
	defer close(pipe)

	go solveSumOfSquares(pipe, param, 1, 0)
	go solveSquareOfSums(pipe, param, 1, 0)

	v := (<-pipe - <-pipe)
	if v < 0 {
		v = v * -1
	}

	fmt.Println("answer ", v)

}
