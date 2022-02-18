package structs

import (
	"reflect"
	"strconv"
)

// IntField returns a slice of int. For more info refer to Slice types IntField() method.
func IntField(s interface{}, fieldName string) []int {
	return NewStructSlice(s).IntField(fieldName)
}

// UintField returns a slice of int. For more info refer to Slice types v() method.
func UintField(s interface{}, fieldName string) []uint {
	return NewStructSlice(s).UintField(fieldName)
}

// Int64Field returns a slice of int64. For more info refer to Slice types Int64Field() method.
func Int64Field(s interface{}, fieldName string) []int64 {
	return NewStructSlice(s).Int64Field(fieldName)
}

// Uint64Field returns a slice of int64. For more info refer to Slice types Uint64Field() method.
func Uint64Field(s interface{}, fieldName string) []uint64 {
	return NewStructSlice(s).Uint64Field(fieldName)
}

// StringField returns a slice of int64. For more info refer to Slice types StringField() method.
func StringField(s interface{}, fieldName string) []string {
	return NewStructSlice(s).StringField(fieldName)
}

// Int returns a slice of int. For more info refer to Slice types Int() method.
func Int(s interface{}) []int {
	return NewStructSlice(s).IntSlice()
}

// Uint returns a slice of uint. For more info refer to Slice types Uint() method.
func Uint(s interface{}) []uint {
	return NewStructSlice(s).Uint()
}

// Int8 returns a slice of int8. For more info refer to Slice types Int8() method.
func Int8(s interface{}) []int8 {
	return NewStructSlice(s).Int8()
}

// Uint8 returns a slice of uint8. For more info refer to Slice types Uint8() method.
func Uint8(s interface{}) []uint8 {
	return NewStructSlice(s).Uint8()
}

// Int16 returns a slice of int16. For more info refer to Slice types Int16() method.
func Int16(s interface{}) []int16 {
	return NewStructSlice(s).Int16()
}

// Uint16 returns a slice of uint16. For more info refer to Slice types Uint16() method.
func Uint16(s interface{}) []uint16 {
	return NewStructSlice(s).Uint16()
}

// Int32 returns a slice of int32. For more info refer to Slice types Int32() method.
func Int32(s interface{}) []int32 {
	return NewStructSlice(s).Int32()
}

// Uint32 returns a slice of uint32. For more info refer to Slice types Uint32() method.
func Uint32(s interface{}) []uint32 {
	return NewStructSlice(s).Uint32()
}

// Int64 returns a slice of int64. For more info refer to Slice types Int64() method.
func Int64(s interface{}) []int64 {
	return NewStructSlice(s).Int64()
}

// Uint64 returns a slice of uint64. For more info refer to Slice types Uint64() method.
func Uint64(s interface{}) []uint64 {
	return NewStructSlice(s).Uint64()
}

// String returns a slice of uint64. For more info refer to Slice types String() method.
func String(s interface{}) []string {
	return NewStructSlice(s).String()
}

// StructSlice hold a struct slice reflect.value
type StructSlice struct {
	value reflect.Value
}

// NewStructSlice returns a new *Slice with the slice s. It panics if the s's kind is not slice.
func NewStructSlice(s interface{}) *StructSlice {
	v := reflect.Indirect(reflect.ValueOf(s))

	if kind := v.Kind(); !(kind == reflect.Slice || kind == reflect.Array) {
		panic("NewStructSlice: require a slice or array")
	}
	return &StructSlice{v}
}

// IntField extracts the given s slice's every element, which is struct, to []int by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
func (s *StructSlice) IntField(fieldName string) []int {
	length := s.value.Len()
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		v := s.structFieldVal(i, fieldName)
		slice[i] = int(valueInteger(v))
	}

	return slice
}

// UintField extracts the given s slice's every element, which is struct, to []uint by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
func (s *StructSlice) UintField(fieldName string) []uint {
	length := s.value.Len()
	slice := make([]uint, length)

	for i := 0; i < length; i++ {
		v := s.structFieldVal(i, fieldName)
		slice[i] = uint(valueInteger(v))
	}

	return slice
}

// Int64Field extracts the given s slice's every element, which is struct, to []int64 by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
func (s *StructSlice) Int64Field(fieldName string) []int64 {
	length := s.value.Len()
	slice := make([]int64, length)

	for i := 0; i < length; i++ {
		v := s.structFieldVal(i, fieldName)
		slice[i] = int64(valueInteger(v))
	}

	return slice
}

// Uint64Field extracts the given s slice's every element, which is struct, to []int64 by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer.
func (s *StructSlice) Uint64Field(fieldName string) []uint64 {
	length := s.value.Len()
	slice := make([]uint64, length)

	for i := 0; i < length; i++ {
		v := s.structFieldVal(i, fieldName)
		slice[i] = valueInteger(v)
	}

	return slice
}

// StringField extracts the given s slice's every element, which is struct, to []string by the field.
// It panics if the s's element is not struct, or field is not exits, or the value of field is not integer or string.
func (s *StructSlice) StringField(fieldName string) []string {
	length := s.value.Len()
	slice := make([]string, length)

	for i := 0; i < length; i++ {
		v := s.structFieldVal(i, fieldName)
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = strconv.FormatUint(v.Uint(), 10)
		case reflect.String:
			slice[i] = v.String()
		case reflect.Float32:
			slice[i] = strconv.FormatFloat(v.Float(), 'f', -1, 32)
		case reflect.Float64:
			slice[i] = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		default:
			panic("StringField: the value of field is not integer or float or string.")
		}
	}
	return slice
}

// Int extracts the given s slice's every element, which is integer or float, to []int by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) IntSlice() []int {
	length := s.value.Len()
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int(v.Uint())
		default:
			panic("Int: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint extracts the given s slice's every element, which is integer or float, to []uint by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Uint() []uint {
	length := s.value.Len()
	slice := make([]uint, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint(v.Uint())
		default:
			panic("Uint: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int8 extracts the given s slice's every element, which is integer or float, to []int8 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Int8() []int8 {
	length := s.value.Len()
	slice := make([]int8, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int8(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int8(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int8(v.Uint())
		default:
			panic("Int8: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint8 extracts the given s slice's every element, which is integer or float, to []uint8 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Uint8() []uint8 {
	length := s.value.Len()
	slice := make([]uint8, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint8(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint8(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint8(v.Uint())
		default:
			panic("Uint8: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int16 extracts the given s slice's every element, which is integer or float, to []int16 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Int16() []int16 {
	length := s.value.Len()
	slice := make([]int16, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int16(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int16(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int16(v.Uint())
		default:
			panic("Int16: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint16 extracts the given s slice's every element, which is integer or float, to []uint16 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Uint16() []uint16 {
	length := s.value.Len()
	slice := make([]uint16, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint16(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint16(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint16(v.Uint())
		default:
			panic("Uint16: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int32 extracts the given s slice's every element, which is integer or float, to []int32 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Int32() []int32 {
	length := s.value.Len()
	slice := make([]int32, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = int32(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = int32(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int32(v.Uint())
		default:
			panic("Int32: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint32 extracts the given s slice's every element, which is integer or float, to []uint32 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Uint32() []uint32 {
	length := s.value.Len()
	slice := make([]uint32, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint32(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint32(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = uint32(v.Uint())
		default:
			panic("Uint32: the value of field is not integer or float.")
		}
	}
	return slice
}

// Int64 extracts the given s slice's every element, which is integer or float, to []int64 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Int64() []int64 {
	length := s.value.Len()
	slice := make([]int64, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = v.Int()
		case reflect.Float32, reflect.Float64:
			slice[i] = int64(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = int64(v.Uint())
		default:
			panic("Int64: the value of field is not integer or float.")
		}
	}
	return slice
}

// Uint64 extracts the given s slice's every element, which is integer or float, to []uint64 by the field.
// It panics if the s's element is not integer or float, or field is not invalid.
func (s *StructSlice) Uint64() []uint64 {
	length := s.value.Len()
	slice := make([]uint64, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = uint64(v.Int())
		case reflect.Float32, reflect.Float64:
			slice[i] = uint64(v.Float())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = v.Uint()
		default:
			panic("Uint64: the value of field is not integer or float.")
		}
	}
	return slice
}

// String extracts the given s slice's every element, which is integer or float or string, to []string by the field.
// It panics if the s's element is not integer or float, string, or field is not invalid.
func (s *StructSlice) String() []string {
	length := s.value.Len()
	slice := make([]string, length)

	for i := 0; i < length; i++ {
		v := reflect.Indirect(s.value.Index(i))
		switch v.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			slice[i] = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			slice[i] = strconv.FormatUint(v.Uint(), 10)
		case reflect.String:
			slice[i] = v.String()
		case reflect.Float32:
			slice[i] = strconv.FormatFloat(v.Float(), 'f', -1, 32)
		case reflect.Float64:
			slice[i] = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		default:
			panic("String: the value of field is not integer or float or string.")
		}
	}
	return slice
}

func (s *StructSlice) structFieldVal(i int, fieldName string) reflect.Value {
	val := s.value.Index(i)
	val = reflect.Indirect(val)

	// check is struct
	if !(val.Kind() != reflect.Invalid && val.Kind() == reflect.Struct) {
		panic("structFieldVal: the slice's element is not struct or pointer of struct!")
	}

	v := val.FieldByName(fieldName)
	if !v.IsValid() {
		panic("structFieldVal: the struct of slice's element has not the field:" + fieldName)
	}
	return v
}

// Name returns the slice's type name within its package. For more info refer
// to Name() function.
func (s *StructSlice) Name() string {
	return s.value.Type().Name()
}

func valueInteger(v reflect.Value) uint64 {
	switch v.Kind() { // nolint: exhaustive
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(v.Int())
	case reflect.Float32, reflect.Float64:
		return uint64(v.Float())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint()
	default:
		panic("valueInteger: the value of field is not integer or float.")
	}
}
