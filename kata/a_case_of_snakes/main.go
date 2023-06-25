package main

import (
	"fmt"
	"strings"
	"unicode"

	cli "github.com/superchausette/hackattic-golang/common"
)

type void struct{}

var types = map[string]void{
	"b":   void{},
	"ch":  void{},
	"d":   void{},
	"dw":  void{},
	"f":   void{},
	"i":   void{},
	"i16": void{},
	"i32": void{},
	"i64": void{},
	"p":   void{},
	"sz":  void{},
	"u16": void{},
	"u32": void{},
	"u64": void{},
	"w":   void{},
}

func isType(word string) bool {
	_, err := types[word]
	return err
}

func ToSnakeCase(input string) string {
	inputAsRune := []rune(input)
	var words []string
	index := 0
	currentWord := ""
	for index < len(input) {
		if unicode.IsUpper(inputAsRune[index]) {
			words = append(words, strings.ToLower(currentWord))
			currentWord = ""
		}
		currentWord += string(inputAsRune[index])
		index++
	}
	words = append(words, strings.ToLower(currentWord))
	if isType(words[0]) {
		words = words[1:]
	}
	return strings.Join(words, "_")
}

func main() {
	input := cli.RetrieveStdin()
	for _, value := range input {
		fmt.Println(ToSnakeCase(value))
	}
}
