package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestImport(t *testing.T) {
	tests := []struct {
		Input  []string
		Output []string
	}{
		{
			Input: []string{
				`{"Bentley.G":{"balance":2134,"account_no":233831255}}`,
				`{"Barclay.E":{"balance":1123,"account_no":312333321}}`,
				`{"Alton.K":{"balance":9315,"account_no":203123613,"extra":{"balance":131}}}`,
				`{"Bancroft.M":{"balance": 233,"account_no":287655771101,"extra":{"balance":98}}}`,
			},
			Output: []string{
				"Bancroft.M: 98",
				"Alton.K: 131",
				"Barclay.E: 1,123",
				"Bentley.G: 2,134",
			},
		},
		{
			Input: []string{
				`{"Allard": {"balance": 5459160, "account_number": 4030714}}`,
				`{"Averill": {"balance": 2286706, "account_number": 3116236}, "extra": {"balance": 5411046}}`,
				`{"Beldon": {"balance": 9258554, "account_number": 4122595}}`,
				`{"Booker": {"balance": 2137674, "account_number": 1373604}}`,
				`{"Bentley": {"balance": 3083480, "account_number": 3766405}}`,
				`{"Brock": {"balance": 5582651, "account_number": 3990369}}`,
				`{"Bradley": {"balance": 9550568, "account_number": 1640425}, "extra": {"balance": 1119993}}`,
				`{"Brinley": {"balance": 6738544, "account_number": 4028239}}`,
				`{"Arundel": {"balance": 3689328, "account_number": 1712599}}`,
				`{"Awarnach": {"balance": 6272168, "account_number": 2979926}}`,
				`{"Bancroft": {"balance": 909540, "account_number": 4565464}}`,
				`{"Burne": {"balance": 4001633, "account_number": 4293397}, "extra": {"balance": 7497821}}`,
				`{"Allston": {"balance": 1722866, "account_number": 2009119}}`,
			},
			Output: []string{
				"Bancroft: 909,540",
				"Bradley: 1,119,993",
				"Allston: 1,722,866",
				"Booker: 2,137,674",
				"Bentley: 3,083,480",
				"Arundel: 3,689,328",
				"Averill: 5,411,046",
				"Allard: 5,459,160",
				"Brock: 5,582,651",
				"Awarnach: 6,272,168",
				"Brinley: 6,738,544",
				"Burne: 7,497,821",
				"Beldon: 9,258,554",
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint("test #", i), func(t *testing.T) {
			output := SortLines(tt.Input)
			if !reflect.DeepEqual(output, tt.Output) {
				t.Fatalf("SortLines(%v) unexpected value got: '%v', want: '%v'", tt.Input, output, tt.Output)
			}
		})
	}
}
