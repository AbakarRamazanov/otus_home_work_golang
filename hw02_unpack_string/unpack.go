package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func createLetters(letter rune, count int) string {
	return strings.Repeat(string(letter), count)
}

func Unpack(inputString string) (string, error) {
	inputRunes := []rune(inputString)
	var result strings.Builder
	step := 1
	count := 1
	i := 0
	for i < len(inputRunes)-1 {
		step = 1
		count = 1
		if unicode.IsDigit(inputRunes[i]) {
			return "", ErrInvalidString
		}
		if inputRunes[i] == '\\' {
			i++
		}
		if (i < len(inputRunes)-1) && unicode.IsDigit(inputRunes[i+1]) {
			count, _ = strconv.Atoi(string(inputRunes[i+1 : i+2]))
			step = 2
		}
		result.WriteString(createLetters(inputRunes[i], count))
		i += step
	}
	if i < len(inputRunes) {
		if !unicode.IsDigit(inputRunes[len(inputRunes)-1]) {
			result.WriteString(createLetters(inputRunes[len(inputRunes)-1], 1))
		} else {
			return "", ErrInvalidString
		}
	}
	return result.String(), nil
}
