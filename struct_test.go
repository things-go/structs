package structs

import (
	"reflect"
	"testing"
)

type UnExportedAndIgnore struct {
	unexported string
	Ignore     int `map:"-"`
}

type NotSupportToString struct {
	ToString []byte `map:"to_string,string"`
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

type EmbedStruct struct {
	EmbedID        int    `map:"embed_id"`
	EmbedOmitValue string `map:"-"`
}

type EmbedPtrStruct struct {
	EmbedID int `map:"embed_id"`
}

type FieldStruct struct {
	ID     int    `map:"id"`
	Ignore string `map:"-"`
}

// var str = "test_string"
// var f func()

func TestEncodeNotAStruct(t *testing.T) {
	var ()

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestEncodeUnexportedAndIgnore(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"unExported and ignore",
			UnExportedAndIgnore{
				unexported: "111",
				Ignore:     111,
			},
			map[string]interface{}{},
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

func TestEncodeToStringNotSupport(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"not a struct",
			NotSupportToString{
				[]byte{1, 2},
			},
			map[string]interface{}{
				"to_string": []byte{1, 2},
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

func TestEncodeBool(t *testing.T) {
	var False bool

	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"bool",
			Bool{
				OmitEmpty:       false,
				NotEmpty:        true,
				NotOmitEmpty:    false,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &False,
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": false, // ignore
				"not_empty":      true,
				"not_omit_empty": false,
				// "omit_empty_ptr":     false, // ignore
				"not_empty_ptr":      false,
				"not_omit_empty_ptr": (*bool)(nil),
			},
		},
		{
			"bool string",
			BoolString{
				OmitEmpty:       false,
				NotEmpty:        true,
				NotOmitEmpty:    false,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &False,
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": false, // ignore
				"not_empty":      "true",
				"not_omit_empty": "false",
				// "omit_empty_ptr":     false, // ignore
				"not_empty_ptr":      "false",
				"not_omit_empty_ptr": (*bool)(nil),
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

func TestEncodeInt(t *testing.T) {
	var (
		IntZero int
		IntOne  = 1
	)
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"int",
			Int{
				OmitEmpty:       IntZero,
				NotEmpty:        IntOne,
				NotOmitEmpty:    IntZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &IntOne,
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": IntZero, // ignore
				"not_empty":      IntOne,
				"not_omit_empty": IntZero,
				// "omit_empty_ptr":     IntZero, // ignore
				"not_empty_ptr":      IntOne,
				"not_omit_empty_ptr": (*int)(nil),
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
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": "0", // ignore
				"not_empty":      "1",
				"not_omit_empty": "0",
				// "omit_empty_ptr":     "0", // ignore
				"not_empty_ptr":      "1",
				"not_omit_empty_ptr": (*int)(nil),
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

func TestEncodeUint(t *testing.T) {
	var (
		UintZero uint
		UintOne  uint = 1
	)
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"uint",
			Uint{
				OmitEmpty:       UintZero,
				NotEmpty:        UintOne,
				NotOmitEmpty:    UintZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &UintOne,
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": UintZero, // ignore
				"not_empty":      UintOne,
				"not_omit_empty": UintZero,
				// "omit_empty_ptr":     UintZero, // ignore
				"not_empty_ptr":      UintOne,
				"not_omit_empty_ptr": (*uint)(nil),
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
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": "0", // ignore
				"not_empty":      "1",
				"not_omit_empty": "0",
				// "omit_empty_ptr":     "0", // ignore
				"not_empty_ptr":      "1",
				"not_omit_empty_ptr": (*uint)(nil),
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

func TestEncodeFloat(t *testing.T) {
	var (
		FloatZero float64
		FloatOne  float64 = 1
	)
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"float",
			Float{
				OmitEmpty:       FloatZero,
				NotEmpty:        FloatOne,
				NotOmitEmpty:    FloatZero,
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &FloatOne,
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": FloatZero, // ignore
				"not_empty":      FloatOne,
				"not_omit_empty": FloatZero,
				// "omit_empty_ptr":     FloatZero, // ignore
				"not_empty_ptr":      FloatOne,
				"not_omit_empty_ptr": (*float64)(nil),
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
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": "0", // ignore
				"not_empty":      "1",
				"not_omit_empty": "0",
				// "omit_empty_ptr":     "0", // ignore
				"not_empty_ptr":      "1",
				"not_omit_empty_ptr": (*float64)(nil),
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

func TestEncodeString(t *testing.T) {
	var StringZero string

	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"string",
			String{
				OmitEmpty:       "",
				NotEmpty:        "1",
				NotOmitEmpty:    "",
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &StringZero,
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": "", // ignore
				"not_empty":      "1",
				"not_omit_empty": "",
				// "omit_empty_ptr":     "", // ignore
				"not_empty_ptr":      "",
				"not_omit_empty_ptr": (*string)(nil),
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

func TestEncodeSlice(t *testing.T) {
	var (
		SliceZero = []byte{}
		SliceOne  = []byte{}
	)
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
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
		{
			"slice struct",
			struct {
				OmitEmpty       []FieldStruct  `map:"omit_empty,omitempty"` // ignore
				NotEmpty        []FieldStruct  `map:"not_empty,omitempty"`
				NotOmitEmpty    []FieldStruct  `map:"not_omit_empty"`
				OmitEmptyPtr    *[]FieldStruct `map:"omit_empty_ptr,omitempty"` // ignore
				NotEmptyPtr     *[]FieldStruct `map:"not_empty_ptr,omitempty"`
				NotOmitEmptyPtr *[]FieldStruct `map:"not_omit_empty_ptr"`
			}{
				OmitEmpty:    []FieldStruct{},
				NotEmpty:     []FieldStruct{{ID: 1}, {ID: 2}},
				OmitEmptyPtr: nil,
				NotEmptyPtr:  &[]FieldStruct{{ID: 1}, {ID: 2}},
				// NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				// "omit_empty": "", // ignore
				"not_empty": []interface{}{
					map[string]interface{}{"id": 1},
					map[string]interface{}{"id": 2},
				},
				"not_omit_empty": []FieldStruct(nil), // TODO: struct空值问题
				// "omit_empty_ptr":     "", // ignore
				"not_empty_ptr": []interface{}{
					map[string]interface{}{"id": 1},
					map[string]interface{}{"id": 2},
				},
				"not_omit_empty_ptr": (*[]FieldStruct)(nil),
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

func TestEncodeEmbedStruct(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"embed struct",
			struct {
				EmbedStruct
			}{
				EmbedStruct: EmbedStruct{
					EmbedID: 111,
				},
			},
			map[string]interface{}{
				"embed_id": 111,
			},
		},
		{
			"embed struct ptr",
			struct {
				*EmbedPtrStruct
			}{
				EmbedPtrStruct: &EmbedPtrStruct{
					EmbedID: 111,
				},
			},
			map[string]interface{}{
				"embed_id": 111,
			},
		},
		{
			"embed struct ptr but nil",
			struct {
				*EmbedPtrStruct
			}{},
			map[string]interface{}{},
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

func TestEncodeFieldStruct(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
		want map[string]interface{}
	}{
		{
			"embed struct",
			struct {
				Field           FieldStruct  `map:"field"`
				OmitEmptyPtr    *FieldStruct `map:"omit_empty_ptr,omitempty"`
				NotEmptyPtr     *FieldStruct `map:"not_empty_ptr"`
				NotOmitEmptyPtr *FieldStruct `map:"not_omit_empty_ptr"`
			}{
				Field:           FieldStruct{ID: 111},
				OmitEmptyPtr:    nil,
				NotEmptyPtr:     &FieldStruct{ID: 222},
				NotOmitEmptyPtr: nil,
			},
			map[string]interface{}{
				"field":              map[string]interface{}{"id": 111},
				"not_empty_ptr":      map[string]interface{}{"id": 222},
				"not_omit_empty_ptr": (*FieldStruct)(nil),
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

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(&IntString{})
	}
}
