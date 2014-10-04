package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(max int) (int, int, int, int) {

	// Using Euclid's formula
	// a = n^2 - m^2
	// b = 2mn
	// c = n^2 + m ^2

	// and our problem states that:
	// a + b + c = 1000

	// substituting Euclid's formula:
	// 2n ^ 2 + 2mn = 1000
	// simply, n ^ 2 + mn = 500

	// let n < square root of 500

	limit := float64(max / 2)
	n := math.Floor(math.Sqrt(limit))
	var m float64

	for n > 0 {
		found := false

		for m = 1; m < n; m++ {
			sum := n*n + m*n
			if sum > limit {
				break
			} else if sum == limit {
				found = true
				break
			}
		}

		if found {
			break
		}

		n--
	}

	a := int(math.Pow(n, 2) - math.Pow(m, 2))
	b := int(2 * m * n)
	c := int(math.Pow(n, 2) + math.Pow(m, 2))

	product := int(a * b * c)

	return a, b, c, product

}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("\nPrints product of a, b, c where a + b + c = NUM and a ^ 2 + b ^ 2 = c ^ 2.\nUsage: go run main.go NUM\n")
		return
	}

	param, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("\nPlease input numeric value only.\n")
		return
	}

	a, b, c, product := solve(param)

	fmt.Printf("a=%v b=%v c=%v, product=%v\n", a, b, c, product)
}
