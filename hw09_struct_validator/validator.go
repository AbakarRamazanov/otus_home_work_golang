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
	// var err error
	for i := 0; i < value.NumField(); i++ {
		tag := value.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}
		switch value.Field(i).Kind() {
		case reflect.String:
			err := validateString(value.Field(i).String(), tag)
			if err != nil {
				if isValidatingError(err) {
					validationErrors = append(validationErrors,
						ValidationError{
							Field: fmt.Sprintf(`%s is %s`, value.Type().Field(i).Name, value.Field(i).String()),
							Err:   errors.Unwrap(err),
						},
					)
				} else {
					return err
				}
			}
		case reflect.Int:
			err := validateInt(value.Field(i).Int(), tag)
			if err != nil {
				if isValidatingError(err) {
					validationErrors = append(validationErrors,
						ValidationError{
							Field: fmt.Sprintf(`%s is %d`, value.Type().Field(i).Name, value.Field(i).Int()),
							Err:   errors.Unwrap(err),
						},
					)
				} else {
					return err
				}
			}
		case reflect.Slice:
			err := validateSlice(value.Field(i), tag)
			if err != nil {
				if isValidatingError(err) {
					validationErrors = append(validationErrors,
						ValidationError{
							Field: fmt.Sprintf(`%s is %v`, value.Type().Field(i).Name, value.Field(i).Slice(0, value.Field(i).Len())),
							Err:   errors.Unwrap(err),
						},
					)
				} else {
					return err
				}
			}
		case reflect.Array, reflect.Bool, reflect.Chan, reflect.Complex128, reflect.Complex64, reflect.Float32,
			reflect.Float64, reflect.Func, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Interface,
			reflect.Invalid, reflect.Map, reflect.Ptr, reflect.Struct, reflect.Uint, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uintptr, reflect.UnsafePointer:
		}
	}
	return validationErrors
}
