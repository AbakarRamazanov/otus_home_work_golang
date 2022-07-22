package hw09structvalidator

import (
	"errors"
	"strings"
)

var (
	ErrorTypeIsNotStruct        = errors.New("type is not struct")
	ErrorValidatingError        = errors.New("validating error")
	ErrorValidatorIsNotValid    = errors.New("validator is not valid")
	ErrorIntLessThanMin         = errors.New("number is less than min")
	ErrorIntMoreThanMax         = errors.New("number is more than max")
	ErrorIntNotIncludedInSet    = errors.New("number is not included in the set")
	ErrorStringLengthIsNotEqual = errors.New("length is not equal")
	ErrorStringRegexpNotMatch   = errors.New("regexp not match string")
	ErrorStringNotIncludedInSet = errors.New("string is not included in the set")
)

func getTags(tag string) []string {
	return strings.Split(tag, "|")
}

func getValidators(validator string) (string, string, error) {
	s := strings.SplitN(validator, ":", 2)
	if len(s) < 2 {
		return "", "", ErrorValidatorIsNotValid
	}
	operand := s[0]
	data := s[1]
	return operand, data, nil
}

func isValidatingError(err error) bool {
	s := strings.SplitN(err.Error(), ":", 2)
	return len(s) == 2 && s[0] == ErrorValidatingError.Error()
}
