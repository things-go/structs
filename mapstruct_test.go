package mapstruct

import (
	"encoding/json"
	"reflect"
	"testing"
)

type ToString struct {
	Number json.Number
	Bool   bool    `map:",string"`
	Int    int     `map:",string"`
	Uint   uint    `map:",string"`
	Float  float64 `map:",string"`
	Array  []byte  `map:",string"`
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

type Ignore struct {
	EmptyBool      bool                   `map:"empty_bool,omitempty"`
	EmptyInt       int                    `map:"empty_int,omitempty"`
	EmptyUint      uint                   `map:"empty_int,omitempty"`
	EmptyFloat     float64                `map:"empty_float,omitempty"`
	EmptyArray     [0]byte                `map:"empty_array,omitempty"`
	EmptySlice     []byte                 `map:"empty_slice,omitempty"`
	EmptyMap       map[string]interface{} `map:"empty_map,omitempty"`
	EmptyString    string                 `map:"empty_string,omitempty"`
	EmptyInterface interface{}            `map:"empty_interface,omitempty"`
	EmptyPtr       *int                   `map:"empty_ptr,omitempty"`
	EmptyPtrStruct *FieldStruct           `map:"empty_ptr_struct,omitempty"`
	IgnoreInt      int                    `map:"-"`
	IgnoreString   string                 `map:"-"`
	IgnoreBool     bool                   `map:"-"`
	Func           func()                 `map:"func,omitempty"`
}

type Exist struct {
	NoEmptyBool       bool    `map:"no_empty_bool"`
	NoEmptyInt        int     `map:"no_empty_int"`
	NoEmptyString     string  `map:"no_empty_string"`
	NoEmptyPtr        *string `map:"no_empty_ptr"`
	NoOmitEmptyBool   bool    `map:"no_omit_empty_bool"`
	NoOmitEmptyInt    int     `map:"no_omit_empty_int"`
	NoOmitEmptyString string  `map:"no_omit_empty_string"`
	NoOmitEmptyPtr    *string `map:"no_omit_empty_ptr"`
	NoEmptyIntString  int     `map:"no_empty_int_string,omitempty,string"`
	EmptyIntString    int     `map:"empty_int_string,omitempty,string"`
	unexported        string
	Field1            FieldStruct  `map:"field1"`
	FieldPtrStruct    *FieldStruct `map:"field_ptr_struct"`
	EmbedStruct
	*EmbedPtrStruct
}

var str = "test_string"
var f func()

func TestEncode(t *testing.T) {
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
			"omitempty",
			Ignore{
				IgnoreInt:    1000,
				IgnoreString: "IgnoreString",
				IgnoreBool:   true,
			},
			map[string]interface{}{"func": f},
		},
		{
			"exist",
			Exist{
				true, 1, "no_empty_string", &str,
				false, 0, "", nil,
				100, 0,
				"",
				FieldStruct{
					111,
					"FieldOmitValue",
				},
				&FieldStruct{
					222,
					"FieldOmitValue",
				},
				EmbedStruct{333, "EmbedOmitValue"},
				&EmbedPtrStruct{"EmbedPtrStruct"},
			},
			map[string]interface{}{
				"no_empty_bool":        true,
				"no_empty_int":         int(1),
				"no_empty_string":      "no_empty_string",
				"no_empty_ptr":         &str,
				"no_omit_empty_bool":   false,
				"no_omit_empty_int":    int(0),
				"no_omit_empty_string": "",
				"no_omit_empty_ptr":    (*string)(nil),
				"no_empty_int_string":  "100",
				"field1":               map[string]interface{}{"field_id": 111},
				"field_ptr_struct":     map[string]interface{}{"field_id": 222},
				"embed_id":             333,
				"embed_ptr_name":       "EmbedPtrStruct",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
