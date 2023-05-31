package main

import (
	"fmt"
	"strconv"
	"time"

	cli "github.com/superchausette/hackattic-golang/common"
)

func WhatDayWasIt(input uint) string {
	t := time.Unix(int64(input)*24*60*60, 0)
	return t.Weekday().String()
}

func main() {
	input := cli.RetrieveStdin()
	for _, line := range input {
		day, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		fmt.Println(WhatDayWasIt(uint(day)))
	}
}
