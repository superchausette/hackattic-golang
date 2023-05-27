package main

import (
	"fmt"

	cli "github.com/superchausette/hackattic-golang/common"
)

const base64Char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func CreateDecodeMap() map[rune]byte {
	decodeMap := make(map[rune]byte)
	for i, c := range base64Char {
		decodeMap[c] = byte(i)
	}
	return decodeMap
}

func Decode4Base64(input []rune, decodeMap map[rune]byte) [4]uint8 {
	if len(input) != 4 {
		panic("Array of length 4 expected")
	}
	var ret [4]uint8
	for i := 0; i < 4; i++ {
		value, ok := decodeMap[input[i]]
		if !ok {
			if input[i] != '=' {
				panic(fmt.Sprint("Invalid base 64 char: ", input[i]))
			}
			ret[i] = 0
		}
		ret[i] = value
	}
	return ret
}

func Debase64(input string, decodeMap map[rune]byte) string {
	var ret string
	inputArray := []rune(input)
	len := len(inputArray)
	idx := 0
	for idx < len {
		switch {
		case len-idx >= 4:
			d := Decode4Base64(inputArray[idx:idx+4], decodeMap)
			b0 := (d[0] << 2) | ((d[1] >> 4) & 0b11)
			b1 := ((d[1] & 0b1111) << 4) | (d[2]>>2)&0b1111
			b2 := ((d[2] & 0b11) << 6) | (d[3])&0b111111
			appendIfNotNull := func(value byte, ret *string) {
				if value == 0 {
					return
				}
				*ret += string(value)
			}
			appendIfNotNull(b0, &ret)
			appendIfNotNull(b1, &ret)
			appendIfNotNull(b2, &ret)
			idx += 4
		default:
			panic(fmt.Sprint("Only ", len, " char remaining"))
		}
	}
	return ret
}

func main() {
	input := cli.RetrieveStdin()
	decodeMap := CreateDecodeMap()
	for _, base64Entry := range input {
		fmt.Println(string(Debase64(base64Entry, decodeMap)))
	}
}
