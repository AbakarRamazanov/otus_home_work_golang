package hw09structvalidator

import (
	"fmt"
	"strconv"
	"strings"
)

func validateInt(value int64, tag string) error {
	tags := getTags(tag)
	for _, t := range tags {
		operand, data, err := getValidators(t)
		if err != nil {
			return err
		}
		switch operand {
		case "min":
			err = intCheckMin(value, data)
			if err != nil {
				return err
			}
		case "max":
			err = intCheckMax(value, data)
			if err != nil {
				return err
			}
		case "in":
			err = intCheckIn(value, data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func intCheckMin(value int64, minString string) error {
	min, err := strconv.Atoi(minString)
	if err != nil {
		return err
	}
	if value < int64(min) {
		return fmt.Errorf("validating error: %w", ErrorIntLessThanMin)
	}
	return nil
}

func intCheckMax(value int64, maxString string) error {
	max, err := strconv.Atoi(maxString)
	if err != nil {
		return err
	}
	if value > int64(max) {
		return fmt.Errorf("validating error: %w", ErrorIntMoreThanMax)
	}
	return nil
}

func intCheckIn(value int64, stringSet string) error {
	sets := strings.Split(stringSet, ",")
	for _, s := range sets {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		if value == int64(i) {
			return nil
		}
	}
	return fmt.Errorf("validating error: %w", ErrorIntNotIncludedInSet)
}
