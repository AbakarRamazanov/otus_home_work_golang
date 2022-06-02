package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type letterCount struct {
	letters []rune
	count   int
}

func Unpack(inputString string) (string, error) {
	inputRunes := []rune(inputString)
	l, err := splitString(inputRunes)
	if err != nil {
		return "", err
	}
	return generateString(l), nil
}

func splitString(inputRune []rune) ([]letterCount, error) {
	l := make([]letterCount, 0)
	for i := 0; i < len(inputRune); {
		if inputRune[i] == '\\' {
			withoutBS, err := processingBackslash(inputRune[i:])
			if err != nil {
				return nil, ErrInvalidString
			}
			inputRune = append(inputRune[:i], withoutBS...)
			i++
			continue
		}
		if unicode.IsDigit(inputRune[i]) {
			lC, err := processingDigit(inputRune[:i+1])
			if err != nil {
				return nil, err
			}
			l = append(l, *lC)
			inputRune = inputRune[i+1:]
			i = 0
			continue
		}
		i++
	}
	if len(inputRune) > 0 {
		l = append(l, letterCount{inputRune, 1})
	}
	return l, nil
}

func processingBackslash(input []rune) ([]rune, error) {
	if len(input) < 2 {
		return nil, ErrInvalidString
	}
	if unicode.IsDigit(input[1]) || input[1] == '\\' {
		input = input[1:]
		return input, nil
	}
	return nil, ErrInvalidString
}

func processingDigit(input []rune) (*letterCount, error) {
	if len(input) == 1 {
		return nil, ErrInvalidString
	}
	count, _ := strconv.Atoi(string(input[len(input)-1]))
	return &letterCount{letters: input[:len(input)-1], count: count}, nil
}

func generateString(l []letterCount) string {
	var builder strings.Builder
	for _, r := range l {
		if len(r.letters) > 1 {
			builder.WriteString(string(r.letters[:len(r.letters)-1]))
		}
		if len(r.letters) > 0 {
			builder.WriteString(strings.Repeat(string(r.letters[len(r.letters)-1]), r.count))
		}
	}
	return builder.String()
}
