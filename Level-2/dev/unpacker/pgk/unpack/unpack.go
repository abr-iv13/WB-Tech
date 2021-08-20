package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (r string, err error) {
	if _, err := strconv.Atoi(s); err == nil {
		return r, errors.New("некорректная строка")
	}

	var prev rune
	var escaped bool
	var b strings.Builder
	for _, char := range s {
		if unicode.IsDigit(char) && !escaped {
			m := int(char - '0')
			r := strings.Repeat(string(prev), m-1)
			b.WriteString(r)
		} else {
			escaped = string(char) == "\\" && string(prev) != "\\"
			if !escaped {
				b.WriteRune(char)
			}
			prev = char
		}
	}

	return b.String(), err
}
