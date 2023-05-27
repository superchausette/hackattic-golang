package main

import (
	"fmt"

	cli "github.com/superchausette/hackattic-golang/common"
)

func AreParensProperlyNested(input string) bool {
	balance := 0
	for _, parens := range input {
		switch parens {
		case '(':
			balance += 1
		case ')':
			balance -= 1
		default:
			panic(fmt.Sprint("Unexpected rune: ", parens))
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

func main() {
	input := cli.RetrieveStdin()
	for _, value := range input {
		if AreParensProperlyNested(value) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
