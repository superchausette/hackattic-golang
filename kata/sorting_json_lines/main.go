package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	cli "github.com/superchausette/hackattic-golang/common"
)

type AccountInfo struct {
	Balance int           `json:"balance"`
	Extra   *ExtraDetails `json:"extra,omitempty"`
}

type ExtraDetails struct {
	Balance int `json:"balance"`
}

type RealEntry struct {
	Name    string
	Balance int
}

func ExtractFromJson(input []string) map[string]AccountInfo {
	entries := make(map[string]AccountInfo)
	for _, jsonInput := range input {
		var entry map[string]AccountInfo
		err := json.Unmarshal([]byte(jsonInput), &entry)
		if err != nil {
			panic(err.Error())
		}
		var extraBalance *int
		if value, ok := entry["extra"]; ok {
			extraBalance = &value.Balance
			delete(entry, "extra")
		}
		for name, info := range entry {
			if extraBalance != nil {
				if info.Extra != nil {
					panic("I don't want to deal with that")
				}
				info.Extra = &ExtraDetails{Balance: *extraBalance}
			}
			entries[name] = info
		}
	}
	return entries
}

func FormatBalance(value int) string {
	defaultStr := strconv.Itoa(value)
	length := len(defaultStr)

	if length <= 3 {
		return defaultStr
	}

	ret := ""
	remainder := length % 3

	if remainder != 0 {
		ret = defaultStr[:remainder] + ","
	}

	for i := remainder; i < length; i += 3 {
		ret += defaultStr[i:i+3] + ","
	}
	ret = ret[:len(ret)-1]

	return ret
}

func SortLines(input []string) []string {
	fromJson := ExtractFromJson(input)
	realEntries := make([]RealEntry, 0, len(fromJson))
	for name, info := range fromJson {
		balance := info.Balance
		if info.Extra != nil {
			balance = info.Extra.Balance
		}
		realEntries = append(realEntries, RealEntry{Name: name, Balance: balance})
	}
	sort.Slice(realEntries, func(i, j int) bool {
		return realEntries[i].Balance < realEntries[j].Balance
	})
	ret := make([]string, 0, len(realEntries))
	for _, entry := range realEntries {
		ret = append(ret, fmt.Sprintf("%s: %s", entry.Name, FormatBalance(entry.Balance)))
	}
	return ret
}

func main() {
	input := cli.RetrieveStdin()
	sortedLines := SortLines(input)
	for _, line := range sortedLines {
		fmt.Println(line)
	}
}
