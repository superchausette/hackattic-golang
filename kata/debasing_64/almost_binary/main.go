package main

import (
	"fmt"

	cli "github.com/superchausette/hackattic-golang/common"
)

func DecodeAlmostBinary(input string) uint16 {
	value := uint16(0)
	for _, c := range input {
		value <<= 1
		switch c {
		case '#':
			value |= 1
		case '.':
			value |= 0
		default:
			panic(fmt.Sprint("Invalid char: ", c))
		}
	}
	return value
}

func main() {
	input := cli.RetrieveStdin()
	for _, line := range input {
		fmt.Println(DecodeAlmostBinary(line))
	}
}
