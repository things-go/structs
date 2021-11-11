package mapstruct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type UnExportedAndIgnore struct {
	unexported string
	IgnoreInt  int `map:"-"`
}

type ToString struct {
	Number json.Number
	Bool   bool    `map:",string"`
	Int    int     `map:",string"`
	Uint   uint    `map:",string"`
	Float  float64 `map:",string"`
	Array  []byte  `map:",string"`
}

type Bool struct {
	OmitEmpty       bool  `map:"omit_empty,omitempty"` // ignore
	NotEmpty        bool  `map:"not_empty,omitempty"`
	NotOmitEmpty    bool  `map:"not_omit_empty"`
	OmitEmptyPtr    *bool `map:"omit_empty_ptr,omitempty"` // ignore
	NotEmptyPtr     *bool `map:"not_empty_ptr,omitempty"`
	NotOmitEmptyPtr *bool `map:"not_omit_empty_ptr"`
}
type BoolString struct {
	OmitEmpty       bool  `map:"omit_empty,omitempty,string"` // ignore
	NotEmpty        bool  `map:"not_empty,omitempty,string"`
	NotOmitEmpty    bool  `map:"not_omit_empty,string"`
	OmitEmptyPtr    *bool `map:"omit_empty_ptr,omitempty,string"` // ignore
	NotEmptyPtr     *bool `map:"not_empty_ptr,omitempty,string"`
	NotOmitEmptyPtr *bool `map:"not_omit_empty_ptr,string"`
}

type Int struct {
	OmitEmpty       int  `map:"omit_empty,omitempty"` // ignore
	NotEmpty        int  `map:"not_empty,omitempty"`
	NotOmitEmpty    int  `map:"not_omit_empty"`
	OmitEmptyPtr    *int `map:"omit_empty_ptr,omitempty"` // ignore
	NotEmptyPtr     *int `map:"not_empty_ptr,omitempty"`
	NotOmitEmptyPtr *int `map:"not_omit_empty_ptr"`
}

type IntString struct {
	OmitEmpty       int  `map:"omit_empty,omitempty,string"` // ignore
	NotEmpty        int  `map:"not_empty,omitempty,string"`
	NotOmitEmpty    int  `map:"not_omit_empty,string"`
	OmitEmptyPtr    *int `map:"omit_empty_ptr,omitempty,string"` // ignore
	NotEmptyPtr     *int `map:"not_empty_ptr,omitempty,string"`
	NotOmitEmptyPtr *int `map:"not_omit_empty_ptr,string"`
}
type Uint struct {
	OmitEmpty       uint  `map:"omit_empty,omitempty"` // ignore
	NotEmpty        uint  `map:"not_empty,omitempty"`
	NotOmitEmpty    uint  `map:"not_omit_empty"`
	OmitEmptyPtr    *uint `map:"omit_empty_ptr,omitempty"` // ignore
	NotEmptyPtr     *uint `map:"not_empty_ptr,omitempty"`
	NotOmitEmptyPtr *uint `map:"not_omit_empty_ptr"`
}
type UintString struct {
	OmitEmpty       uint  `map:"omit_empty,omitempty,string"` // ignore
	NotEmpty        uint  `map:"not_empty,omitempty,string"`
	NotOmitEmpty    uint  `map:"not_omit_empty,string"`
	OmitEmptyPtr    *uint `map:"omit_empty_ptr,omitempty,string"` // ignore
	NotEmptyPtr     *uint `map:"not_empty_ptr,omitempty,string"`
	NotOmitEmptyPtr *uint `map:"not_omit_empty_ptr,string"`
}

type Float struct {
	OmitEmpty       float64  `map:"omit_empty,omitempty"` // ignore
	NotEmpty        float64  `map:"not_empty,omitempty"`
	NotOmitEmpty    float64  `map:"not_omit_empty"`
	OmitEmptyPtr    *float64 `map:"omit_empty_ptr,omitempty"` // ignore
	NotEmptyPtr     *float64 `map:"not_empty_ptr,omitempty"`
	NotOmitEmptyPtr *float64 `map:"not_omit_empty_ptr"`
}
type FloatString struct {
	OmitEmpty       float64  `map:"omit_empty,omitempty,string"` // ignore
	NotEmpty        float64  `map:"not_empty,omitempty,string"`
	NotOmitEmpty    float64  `map:"not_omit_empty,string"`
	OmitEmptyPtr    *float64 `map:"omit_empty_ptr,omitempty,string"` // ignore
	NotEmptyPtr     *float64 `map:"not_empty_ptr,omitempty,string"`
	NotOmitEmptyPtr *float64 `map:"not_omit_empty_ptr,string"`
}

type String struct {
	OmitEmpty       string  `map:"omit_empty,omitempty"` // ignore
	NotEmpty        string  `map:"not_empty,omitempty"`
	NotOmitEmpty    string  `map:"not_omit_empty"`
	OmitEmptyPtr    *string `map:"omit_empty_ptr,omitempty"` // ignore
	NotEmptyPtr     *string `map:"not_empty_ptr,omitempty"`
	NotOmitEmptyPtr *string `map:"not_omit_empty_ptr"`
}

type Slice struct {
	OmitEmpty       []byte  `map:"omit_empty,omitempty"` // ignore
	NotEmpty        []byte  `map:"not_empty,omitempty"`
	NotOmitEmpty    []byte  `map:"not_omit_empty"`
	OmitEmptyPtr    *[]byte `map:"omit_empty_ptr,omitempty"` // ignore
	NotEmptyPtr     *[]byte `map:"not_empty_ptr,omitempty"`
	NotOmitEmptyPtr *[]byte `map:"not_omit_empty_ptr"`
}

type Ignore struct {
	EmptyArray     [0]byte                `map:"empty_array,omitempty"`
	EmptyMap       map[string]interface{} `map:"empty_map,omitempty"`
	EmptyInterface interface{}            `map:"empty_interface,omitempty"`
	EmptyPtrStruct *FieldStruct           `map:"empty_ptr_struct,omitempty"`
	Func           func()                 `map:"func,omitempty"`
}

type FieldStruct struct {
	FieldID        int    `map:"field_id"`
	FieldOmitValue string `map:"-"`
}

type EmbedStruct struct {
	EmbedID        int    `map:"embed_id"`
	EmbedOmitValue string `map:"-"`
}

type EmbedPtrStruct struct {
	EmbedPtrName string `map:"embed_ptr_name"`
}

// 	EmbedStruct
//	*EmbedPtrStruct
// Field1            FieldStruct   `map:"field1"`
// FieldPtrStruct    *FieldStruct  `map:"field_ptr_struct"`
// Field2            []FieldStruct `map:"field2,omitempty"`
// Field3            []FieldStruct `map:"field3"`
// Field4            []FieldStruct `map:"field4"`

var str = "test_string"
var f func()

func TestEncode(t *testing.T) {
	var (
		False      bool
		True       = true
		IntZero    int
		IntOne     = 1
		UintZero   uint
		UintOne    uint = 1
		FloatZero  float64
		FloatOne   float64 = 1
		StringZero string
		StringOne  = "1"
		SliceZero  = []byte{}
		SliceOne   = []byte{}
	)

	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"not a struct",
			"not a struct",
			nil,
		},
		{
			"unExported and ignore",
			UnExportedAndIgnore{
				unexported: "111",
				IgnoreInt:  111,
			},
			map[string]interface{}{},
		},
		{
			"to string",
			ToString{
				Number: json.Number("555"),
				Bool:   true,
				Int:    11,
				Uint:   22,
				Float:  3.33,
			},
			map[string]interface{}{
				"Number": json.Number("555"),
				"Bool":   "true",
				"Int":    "11",
				"Uint":   "22",
				"Float":  "3.33",
				"Array":  []byte(nil),
			},
		},
		{
			"bool",
			Bool{
				OmitEmpty:       false,
				NotEmpty:        true,
				NotOmitEmpty:    false,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &True,
				NotOmitEmptyPtr: &False,
			},
			map[string]interface{}{
				// "omit_empty": false, // ignore
				"not_empty":      true,
				"not_omit_empty": false,
				// "omit_empty_ptr":     false, // ignore
				"not_empty_ptr":      true,
				"not_omit_empty_ptr": false,
			},
		},
		{
			"bool string",
			BoolString{
				OmitEmpty:       false,
				NotEmpty:        true,
				NotOmitEmpty:    false,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &True,
				NotOmitEmptyPtr: &False,
			},
			map[string]interface{}{
				// "omit_empty": false, // ignore
				"not_empty":      "true",
				"not_omit_empty": "false",
				// "omit_empty_ptr":     false, // ignore
				"not_empty_ptr":      "true",
				"not_omit_empty_ptr": "false",
			},
		},
		{
			"int",
			Int{
				OmitEmpty:       IntZero,
				NotEmpty:        IntOne,
				NotOmitEmpty:    IntZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &IntOne,
				NotOmitEmptyPtr: &IntZero,
			},
			map[string]interface{}{
				// "omit_empty": IntZero, // ignore
				"not_empty":      IntOne,
				"not_omit_empty": IntZero,
				// "omit_empty_ptr":     IntZero, // ignore
				"not_empty_ptr":      IntOne,
				"not_omit_empty_ptr": IntZero,
			},
		},
		{
			"int string",
			IntString{
				OmitEmpty:       IntZero,
				NotEmpty:        IntOne,
				NotOmitEmpty:    IntZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &IntOne,
				NotOmitEmptyPtr: &IntZero,
			},
			map[string]interface{}{
				// "omit_empty": "0", // ignore
				"not_empty":      "1",
				"not_omit_empty": "0",
				// "omit_empty_ptr":     "0", // ignore
				"not_empty_ptr":      "1",
				"not_omit_empty_ptr": "0",
			},
		},
		{
			"uint",
			Uint{
				OmitEmpty:       UintZero,
				NotEmpty:        UintOne,
				NotOmitEmpty:    UintZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &UintOne,
				NotOmitEmptyPtr: &UintZero,
			},
			map[string]interface{}{
				// "omit_empty": UintZero, // ignore
				"not_empty":      UintOne,
				"not_omit_empty": UintZero,
				// "omit_empty_ptr":     UintZero, // ignore
				"not_empty_ptr":      UintOne,
				"not_omit_empty_ptr": UintZero,
			},
		},
		{
			"uint string",
			UintString{
				OmitEmpty:       UintZero,
				NotEmpty:        UintOne,
				NotOmitEmpty:    UintZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &UintOne,
				NotOmitEmptyPtr: &UintZero,
			},
			map[string]interface{}{
				// "omit_empty": "0", // ignore
				"not_empty":      "1",
				"not_omit_empty": "0",
				// "omit_empty_ptr":     "0", // ignore
				"not_empty_ptr":      "1",
				"not_omit_empty_ptr": "0",
			},
		},
		{
			"float",
			Float{
				OmitEmpty:       FloatZero,
				NotEmpty:        FloatOne,
				NotOmitEmpty:    FloatZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &FloatOne,
				NotOmitEmptyPtr: &FloatZero,
			},
			map[string]interface{}{
				// "omit_empty": FloatZero, // ignore
				"not_empty":      FloatOne,
				"not_omit_empty": FloatZero,
				// "omit_empty_ptr":     FloatZero, // ignore
				"not_empty_ptr":      FloatOne,
				"not_omit_empty_ptr": FloatZero,
			},
		},
		{
			"float string",
			FloatString{
				OmitEmpty:       FloatZero,
				NotEmpty:        FloatOne,
				NotOmitEmpty:    FloatZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &FloatOne,
				NotOmitEmptyPtr: &FloatZero,
			},
			map[string]interface{}{
				// "omit_empty": "0", // ignore
				"not_empty":      "1",
				"not_omit_empty": "0",
				// "omit_empty_ptr":     "0", // ignore
				"not_empty_ptr":      "1",
				"not_omit_empty_ptr": "0",
			},
		},
		{
			"string",
			String{
				OmitEmpty:       "",
				NotEmpty:        "1",
				NotOmitEmpty:    "",
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &StringOne,
				NotOmitEmptyPtr: &StringZero,
			},
			map[string]interface{}{
				// "omit_empty": "", // ignore
				"not_empty":      "1",
				"not_omit_empty": "",
				// "omit_empty_ptr":     "", // ignore
				"not_empty_ptr":      "1",
				"not_omit_empty_ptr": "",
			},
		},
		{
			"slice",
			Slice{
				OmitEmpty:       []byte{},
				NotEmpty:        []byte{1},
				NotOmitEmpty:    []byte{},
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &SliceOne,
				NotOmitEmptyPtr: &SliceZero,
			},
			map[string]interface{}{
				// "omit_empty": "", // ignore
				"not_empty":      []byte{1},
				"not_omit_empty": []byte{},
				// "omit_empty_ptr":     "", // ignore
				"not_empty_ptr":      SliceOne,
				"not_omit_empty_ptr": SliceZero,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

// func BenchmarkEncode(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Encode(&Exist{
// 			true, 1, "no_empty_string", &str,
// 			false, 0, "", nil,
// 			100, 0,
// 			"",
// 			FieldStruct{
// 				111,
// 				"FieldOmitValue",
// 			},
// 			&FieldStruct{
// 				222,
// 				"FieldOmitValue",
// 			},
// 			EmbedStruct{333, "EmbedOmitValue"},
// 			&EmbedPtrStruct{"EmbedPtrStruct"},
// 		})
// 	}
// }

func TestName(t *testing.T) {
	var a [4]byte

	vv := reflect.ValueOf(a)

	switch s := vv.Interface().(type) {
	case []byte:
		fmt.Println(s)
	default:

	}

}
