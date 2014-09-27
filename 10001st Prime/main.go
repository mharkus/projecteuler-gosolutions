package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func solve(out chan int, n int) {
	primes := []int{}

	for i := 2; len(primes) < n; i++ {

		if isPrime(i, primes) {
			primes = append(primes, i)
		}
	}

	out <- primes[n-1]
}

func isPrime(i int, primes []int) bool {

	for _, prime := range primes {
		if i%prime == 0 {
			return false
		}
	}

	return true
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("\nComputes for the NUM'th prime number\nUsage: go run main.go NUM\n")
		return
	}

	param, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("\nPlease input numeric value only.\n")
		return
	}

	out := make(chan int)
	defer close(out)

	start := time.Now()
	go solve(out, param)

	fmt.Printf("%[1]dth prime is %[2]d\n", param, <-out)

	fmt.Println("Total execution time:", time.Since(start))
}
