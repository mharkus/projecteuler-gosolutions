package main

import (
    "fmt"
)

const MAX = 1000

func solve(max, sum int) {
	if(max == 0){
		fmt.Println("the sum is ", sum)
		return
	}

	if max < MAX && (max % 3 == 0 || max % 5 == 0) {
		solve(max - 1, sum + max)
	}else{
		solve(max - 1, sum)	
	}
}

func main() {
    solve(MAX, 0)
}