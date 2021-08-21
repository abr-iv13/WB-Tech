package unpack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {

	testTable := []struct {
		in  string
		out string
	}{

		{
			in:  "a4bc2d5e",
			out: "aaaabccddddde",
		},
		{
			in:  "abcd",
			out: "abcd",
		},
		{
			in:  "45",
			out: "",
		},
		{
			in:  "",
			out: "",
		},
	}

	for _, testCase := range testTable {
		result, err := Unpack(testCase.in)
		assert.Equal(t, testCase.out, result,
			fmt.Sprintf("Incorrected result, Except %s, got %s", testCase.out, result),
		)

		if err != nil {
			assert.Error(t, err)
		}
	}
}
