package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

//Преобразование string в rune, проверка строки на наличии пробелов
func isEscape(char rune, prev rune) bool {
	return string(char) == "\\" && string(prev) != "\\"
}

//Проверка руны на десятичное число и на наличие пробелов
func isDecimalRuneAndEscaped(char rune, escaped bool) bool {
	//IsDigit проверяет является ли руна десятичной цифрой
	return unicode.IsDigit(char) && !escaped
}

func Unpack(s string) (string, error) {
	// Проверка на пустую строку
	if len(s) == 0 {
		return "", nil
	}
	//Преобразование строки в int
	if _, err := strconv.Atoi(s); err == nil {
		return "", errors.New("Некорректная строка")
	}

	var prev rune         //Предыдущее значение
	var escaped bool      //flag на наличие пробелов
	var b strings.Builder //Построитель используется для эффективного построения строки с использованием методов записи.

	for _, char := range s {
		if isDecimalRuneAndEscaped(char, escaped) {
			//Преобразование руны в число
			//«вычитание значения руны« 0 »из любой руны от« 0 »до« 9 »даст целое число от 0 до 9».
			i := int(char - '0')
			//Repeat возвращает новую строку, состоящую из количества копий строки.
			r := strings.Repeat(string(prev), i-1)
			//Добавление строку в буфер
			b.WriteString(r)
		} else {
			escaped = isEscape(char, prev)
			if !escaped {
				//Добавление rune в буфер
				b.WriteRune(char)
			}

			prev = char
		}
	}

	return b.String(), nil
}
