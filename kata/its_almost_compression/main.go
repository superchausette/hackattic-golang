package main

import (
	"fmt"
	"strconv"

	cli "github.com/superchausette/hackattic-golang/common"
)

func Compress(input string) string {
	var lastLetter rune
	lastLetterRepetition := 0
	var compressed string

	onLetterChange := func() {
		if lastLetterRepetition <= 2 {
			compressed += string(lastLetter)
			if lastLetterRepetition == 2 {
				compressed += string(lastLetter)
			}
		} else {
			compressed += strconv.Itoa(lastLetterRepetition) + string(lastLetter)
		}
	}

	for _, currentLetter := range input {
		if lastLetterRepetition == 0 {
			lastLetter = currentLetter
		}
		if currentLetter != lastLetter {
			onLetterChange()
			lastLetter = currentLetter
			lastLetterRepetition = 0
		}
		lastLetterRepetition++
	}
	onLetterChange()
	return compressed
}

func main() {
	input := cli.RetrieveStdin()

	for _, text := range input {
		fmt.Println(Compress(text))
	}
}
