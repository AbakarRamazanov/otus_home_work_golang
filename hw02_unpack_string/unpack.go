package hw02unpackstring

import (
	"errors"
	"fmt"
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
	err := checkBackslash(inputRunes)
	if err != nil {
		return "", err
	}
	l, err := splitString(inputRunes)
	if err != nil {
		return "", err
	}
	deleteBackslash(l)
	resultTwo := generateString(l)
	return resultTwo, nil
}

func splitString(inputRune []rune) ([]letterCount, error) {
	l := make([]letterCount, 0)
	for i := 0; i < len(inputRune); {
		if inputRune[i] == '\\' {
			fmt.Println("before input:", string(inputRune))
			dInputRune, err := checkBackslash2(inputRune[i:])
			if err != nil {
				return nil, ErrInvalidString
			}
			inputRune := append(inputRune[:i], dInputRune...)
			fmt.Println("after input:", string(inputRune))
			i += 2
			continue
		}
		if unicode.IsDigit(inputRune[i]) {
			if i == 0 {
				return nil, ErrInvalidString
			}
			count, _ := strconv.Atoi(string(inputRune[i : i+1]))
			l = append(l, letterCount{letters: inputRune[:i], count: count})
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

func checkBackslash2(input []rune) ([]rune, error) {
	if len(input) < 2 {
		return nil, ErrInvalidString
	}
	if unicode.IsDigit(input[1]) || input[1] == '\\' {
		input = input[1:]
		fmt.Println("checkBackslash2 input:", string(input))
		return input, nil
	}
	return nil, ErrInvalidString
}

func checkBackslash(input []rune) error {
	i := 0
	for ; i < len(input)-1; i++ {
		if input[i] == '\\' {
			if unicode.IsDigit(input[i+1]) || input[i+1] == '\\' {
				i++
			} else {
				return ErrInvalidString
			}
		}
	}
	if i < len(input) && input[i] == '\\' {
		return ErrInvalidString
	}
	return nil
}

func deleteBackslash(l []letterCount) {
	// return
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l[i].letters); j++ {
			if l[i].letters[j] == '\\' {
				l[i].letters = append(l[i].letters[:j], l[i].letters[j+1:]...)
			}
		}
	}
}

func generateString(l []letterCount) string {
	var builder strings.Builder
	for _, r := range l {
		if len(r.letters) > 1 {
			builder.WriteString(string(r.letters[:len(r.letters)-1]))
		}
		if len(r.letters) > 0 {
			builder.WriteString(createLetters(r.letters[len(r.letters)-1], r.count))
		}
	}
	return builder.String()
}

func createLetters(letter rune, count int) string {
	return strings.Repeat(string(letter), count)
}
