// Package mapstruct Go library for encoding native Go structures into generic map values.
//
// The simplest function to start with is Encoded.
//
// Field Tags
//
// When encode to a map[string]interface{}, mapstruct will use the field
// name by default to perform the mapping. For example, if a struct has
// a field "Username" then mapstruct will use a key "Username".
//
//     type User struct {
//         Username string
//     }
//
// You can change the behavior of mapstruct by using struct tags.
// The default struct tag that mapstruct looks for is "map"
// but you can customize it using EncodeWithTag.
//
// Renaming Fields
//
// To rename the key that mapstruct looks for, use the "map" tag and
// set a value directly. For example, to change the "username" example
// above to "user":
//
//     type User struct {
//         Username string `map:"user"`
//     }
//
// Embedded Structs
//
// Embedded structs are treated as if they're another field with that name.
//
//     type Person struct {
//         Name string `map:"name"`
//     }
//
//     type Friend struct {
//         Person
//     }
//
// This would output that looks like below:
//
//     map[string]interface{}{
//         "name": "alice",
//     }
//
// Omit Empty Values
//
// When encoding from a struct to any other value, you may use the
// ",omitempty" suffix on your tag to omit that value if it equates to
// the zero value. The zero value of all types is specified in the Go
// specification.
//
// For example, the zero type of a numeric type is zero ("0"). If the struct
// field value is zero and a numeric type, the field is empty, and it won't
// be encoded into the destination type.
//
//     type Source {
//         Age int `map:",omitempty"`
//     }
//
// Unexported fields
//
// Since unexported (private) struct fields cannot be set outside the package
// where they are defined, the encoder will simply skip them.
//
// For this input type definition:
//
//     type Exported struct {
//         private string // this unexported field will be skipped
//         Public string
//     }
//
// this map as output:
//
//     map[string]interface{}{
//         "Public":  "I made it through!",
//     }
package mapstruct

import (
	"reflect"
	"strconv"
)

const DefaultTag = "map"

// Encode takes an input structure and uses reflection to translate it to
// the output map[string]interface{} with default tag "map"
func Encode(input interface{}) map[string]interface{} { return EncodeWithTag(input, DefaultTag) }

// EncodeWithTag takes an input structure and uses reflection to translate it to
// the output map[string]interface{} with the custom tag name
func EncodeWithTag(input interface{}, tagName string) map[string]interface{} {
	t := reflect.TypeOf(input)
	v := reflect.ValueOf(input)

	if v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	m := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)

		if !fv.CanInterface() || ft.PkgPath != "" {
			continue
		}

		tag := ft.Tag.Get(tagName)
		if tag == "-" { // skip the field
			continue
		}
		keyName := ft.Name
		name, opts := parseTag(tag)
		if isValidTag(name) {
			keyName = name
		}
		if opts.Contains("omitempty") && isEmptyValue(fv) { // skip empty field
			continue
		}

		// ft.Anonymous means embedded field
		if ft.Anonymous {
			if (fv.Kind() == reflect.Struct) ||
				(fv.Kind() == reflect.Ptr && !fv.IsNil() && fv.Elem().Kind() == reflect.Struct) {
				embedded := EncodeWithTag(fv.Interface(), tagName)
				for embName, embValue := range embedded {
					m[embName] = embValue
				}
			}
			continue
		}

		if opts.Contains("string") {
			if str := toString(fv); str != nil {
				m[keyName] = str
				continue
			}
		}
		if (fv.Kind() == reflect.Struct) ||
			(fv.Kind() == reflect.Ptr && fv.Elem().Kind() == reflect.Struct) {
			m[keyName] = EncodeWithTag(fv.Interface(), tagName)
			continue
		}
		m[keyName] = fv.Interface()
	}
	return m
}

func toString(fv reflect.Value) interface{} {
	switch k := fv.Kind(); k {
	case reflect.Bool:
		return strconv.FormatBool(fv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(fv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(fv.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(fv.Float(), 'f', 2, 64)
		// TODO: support other types
	}
	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch k := v.Kind(); k {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Array, reflect.Slice, reflect.Map, reflect.String:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}