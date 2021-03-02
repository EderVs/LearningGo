package main

import (
	"reflect"
	"strconv"
	"strings"
)

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// floating point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var structS strings.Builder
		structS.WriteString("{ ")
		for i := 0; i < v.NumField(); i++ {
			structS.WriteString(v.Type().Field(i).Name)
			structS.WriteString(": ")
			structS.WriteString(formatAtom(v.Field(i)))
			structS.WriteString(", ")
		}
		structS.WriteString("}")
		return structS.String()
	case reflect.Array:
		var structS strings.Builder
		structS.WriteString("[ ")
		for i := 0; i < v.Len(); i++ {
			structS.WriteString(formatAtom(v.Index(i)))
			structS.WriteString(", ")
		}
		structS.WriteString("]")
		return structS.String()
	default: // reflect.Interface
		return v.Type().String() + " value"
	}
}
