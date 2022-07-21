package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var s string
	for i := 0; i < len(v); i++ {
		s = fmt.Sprintf("%sField %s invalid, error: %s; ", s, v[i].Field, v[i].Err.Error())
	}
	return s
}

func Validate(v interface{}) error {
	validationErrors := make(ValidationErrors, 0)
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Struct {
		return ErrorTypeIsNotStruct
	}
	for i := 0; i < value.NumField(); i++ {
		tag := value.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		validationError, err := checkField(value, i, tag)
		if err != nil {
			return err
		}
		if validationError != nil {
			validationErrors = append(validationErrors, *validationError)
		}
	}
	if len(validationErrors) != 0 {
		return validationErrors
	}
	return nil
}

func checkField(value reflect.Value, i int, tag string) (*ValidationError, error) {
	switch value.Field(i).Kind() {
	case reflect.String:
		err := validateString(value.Field(i).String(), tag)
		if err != nil {
			if isValidatingError(err) {
				return makeVE(value.Type().Field(i).Name, value.Field(i).String(), err), nil
			}
			return nil, err
		}
	case reflect.Int:
		err := validateInt(value.Field(i).Int(), tag)
		if err != nil {
			return makeVE(value.Type().Field(i).Name, value.Field(i).Int(), err), nil
		}
		return nil, err
	case reflect.Slice:
		err := validateSlice(value.Field(i), tag)
		if err != nil {
			if isValidatingError(err) {
				return makeVE(value.Type().Field(i).Name, value.Field(i).Slice(0, value.Field(i).Len()), err), nil
			}
			return nil, err
		}
	case reflect.Array, reflect.Bool, reflect.Chan, reflect.Complex128, reflect.Complex64, reflect.Float32,
		reflect.Float64, reflect.Func, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Interface,
		reflect.Invalid, reflect.Map, reflect.Ptr, reflect.Struct, reflect.Uint, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uintptr, reflect.UnsafePointer:
	}
	return nil, nil
}

func makeVE(name string, i interface{}, err error) *ValidationError {
	return &ValidationError{
		Field: fmt.Sprintf(`%s is %v`, name, i),
		Err:   errors.Unwrap(err),
	}
}
