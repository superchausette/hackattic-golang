package cli

import (
	"io"
	"os"
	"strings"
)

func RetrieveStdin() []string {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	nonFilteredStdin := strings.Split(string(stdin), "\n")
	ret := []string{}
	for _, input := range nonFilteredStdin {
		if len(input) > 0 {
			ret = append(ret, input)
		}
	}
	return ret
}
