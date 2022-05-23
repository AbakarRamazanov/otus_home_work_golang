package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

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
		if unicode.IsDigit(inputRunes[i+1]) {
			count, _ = strconv.Atoi(string(inputRunes[i+1 : i+2]))
			step = 2
		}
		result.WriteString(strings.Repeat(string(inputRunes[i]), count))
		i += step
	}
	if i < len(inputRunes) && !unicode.IsDigit(inputRunes[len(inputRunes)-1]) {
		result.WriteString(string(inputRunes[len(inputRunes)-1]))
	}
	return result.String(), nil
}
