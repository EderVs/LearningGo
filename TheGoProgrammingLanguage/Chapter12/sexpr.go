package main

import (
	"bytes"
	"fmt"
	"reflect"
)

// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value, spaces int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Bool:
		fmt.Fprintf(buf, "%v", v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), spaces)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				writeNSpaces(buf, spaces+1)
			}
			if err := encode(buf, v.Index(i), spaces+1); err != nil {
				return err
			}
			buf.WriteString("\n")
		}
		if v.Len() > 0 {
			writeNSpaces(buf, spaces)
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				writeNSpaces(buf, spaces+1)
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), spaces+1+1+len(v.Type().Field(i).Name)+1); err != nil {
				return err
			}
			buf.WriteString(")\n")
		}
		if v.NumField() > 0 {
			writeNSpaces(buf, spaces)
		}
		buf.WriteString(")")

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				writeNSpaces(buf, spaces+1)
			}
			buf.WriteByte('(')
			if err := encode(buf, key, spaces+1+1); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), spaces+1+len(key.String())+1+3); err != nil {
				return err
			}
			buf.WriteString(")\n")
		}
		if len(v.MapKeys()) > 0 {
			writeNSpaces(buf, spaces)
		}
		buf.WriteString(")")

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsopported type: %s", v.Type())
	}
	return nil
}

func writeNSpaces(buf *bytes.Buffer, spaces int) {
	for space := 0; space < spaces; space++ {
		buf.WriteByte(' ')
	}
}
