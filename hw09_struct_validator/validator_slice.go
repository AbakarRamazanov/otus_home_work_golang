package hw09structvalidator

import (
	"reflect"
)

func validateSlice(value reflect.Value, tag string) error {
	if value.Len() == 0 {
		return nil
	}
	switch value.Index(0).Kind() {
	case reflect.String:
		err := validateSliceString(value.Interface().([]string), tag)
		if err != nil {
			return err
		}
	case reflect.Int:
		err := validateSliceInt(value.Interface().([]int), tag)
		if err != nil {
			return err
		}
	case reflect.Array, reflect.Bool, reflect.Chan, reflect.Complex128, reflect.Complex64, reflect.Float32,
		reflect.Float64, reflect.Func, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Interface,
		reflect.Invalid, reflect.Map, reflect.Pointer, reflect.Slice, reflect.Struct, reflect.Uint, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Uintptr, reflect.UnsafePointer:
	}
	return nil
}

func validateSliceString(value []string, tag string) error {
	for i := 0; i < len(value); i++ {
		err := validateString(value[i], tag)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateSliceInt(value []int, tag string) error {
	for i := 0; i < len(value); i++ {
		err := validateInt(int64(value[i]), tag)
		if err != nil {
			return err
		}
	}
	return nil
}
