package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(value int) string {
	if value%3 == 0 && value%5 == 0 {
		return "FizzBuzz"
	}
	if value%3 == 0 {
		return "Fizz"
	}
	if value%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(value)
}

func main() {
	var start int
	var end int
	fmt.Scanf("%d %d", &start, &end)
	for i := start; i <= end; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
