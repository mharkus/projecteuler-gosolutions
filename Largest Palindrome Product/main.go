package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func solve(x, y, largest int) {
	if x < 100 && y < 100 {
		fmt.Printf("largest palindrome\n", largest)
		return
	} else if x < 100 {
		solve(999, y-1, largest)
		return
	}

	product := x * y
	if isPalindrome(product) {

		if largest < product {
			largest = product
		}
	}

	solve(x-1, y, largest)
}

func isPalindrome(num int) bool {
	val := strconv.Itoa(num)
	return val == reverse(val)
}

func reverse(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)

	for _, c := range s {
		i--
		cs[i] = c
	}

	return string(cs)
}

func main() {
	solve(999, 999, 0)
}
