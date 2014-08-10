package main

import(
	"fmt"
)

const UPPER_LIMIT = 4000000

func solve(n1, n2, sum int){

	if(n2 > UPPER_LIMIT){
		fmt.Println("sum of even series is " , sum)
		return
	}

	temp := 0

	if(n2 % 2 == 0){
		temp += n2
	}

	solve(n2, n1 + n2, temp + sum)
}

func main() {
	solve(1, 2, 0)
}