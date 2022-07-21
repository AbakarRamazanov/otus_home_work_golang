package hw09structvalidator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func validateString(value string, tag string) error {
	tags := getTags(tag)
	for _, t := range tags {
		operand, data, err := getValidators(t)
		if err != nil {
			return err
		}
		switch operand {
		case "len":
			l, err := strconv.Atoi(data)
			if err != nil {
				return err
			}
			err = stringCheckLength(value, l)
			if err != nil {
				return err
			}
		case "regexp":
			r, err := regexp.Compile(data)
			if err != nil {
				return err
			}
			err = stringCheckRegexp(value, r)
			if err != nil {
				return err
			}
		case "in":
			err := stringCheckIn(value, data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func stringCheckLength(value string, length int) error {
	if utf8.RuneCountInString(value) != length {
		return fmt.Errorf("validating error: %w", ErrorStringLengthIsNotEqual)
	}
	return nil
}

func stringCheckRegexp(value string, r *regexp.Regexp) error {
	if !r.MatchString(value) {
		return fmt.Errorf("%v: %w", ErrorValidatingError, ErrorStringRegexpNotMatch)
	}
	return nil
}

func stringCheckIn(value string, stringSet string) error {
	sets := strings.Split(stringSet, ",")
	for _, s := range sets {
		if value == s {
			return nil
		}
	}
	return fmt.Errorf("validating error: %w", ErrorStringNotIncludedInSet)
}
