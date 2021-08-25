package anagram

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSetsAnagram(t *testing.T) {
	testTable  := []struct{
		//name string
		in []string
		out map[string][]string
	} {
		{
			in:  []string{"тяпка", "пятка", "пятак"},
			out: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}},
		},
		{
			in:  []string{"тЯпкА", "ПЯТКА", "пЯтАК"},
			out: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}},
		},
		{
			in: []string{"тяпка", "пятка", "пятак", "тяпка", "пятка", "пятак"},
			out: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}},
		},
	}

	for _, testCase:=  range testTable {
		t.Run("testCase", func(t *testing.T) {
			result := FindSetsAnagram(testCase.in)
			assert.Equal(t, testCase.out, result,
				fmt.Sprintf("invalid result exept = %s, result = $%s ", testCase.out, result))

		})
	}
}
