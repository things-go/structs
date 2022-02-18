package structs

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type StructInt8 struct {
	UID   int8
	Value string
}

type StructUint16 struct {
	UID   uint16
	Value string
}

type StructFloat64 struct {
	UID   float64
	Value string
}

type StructFloat32 struct {
	UID   float32
	Value string
}

type StructString struct {
	UID   string
	Value string
}

type StructMuch struct {
	UID *string
	Err error
}

func TestStructsIntSlice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"array no ptr",
			args{
				[2]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
		{
			"array ptr",
			args{
				[2]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntField(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructsUintSlice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"array no ptr",
			args{
				[2]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
		{
			"array ptr",
			args{
				[2]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintField(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructsInt64Slice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"array no ptr",
			args{
				[2]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
		{
			"array ptr",
			args{
				[2]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]int64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Field(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Field() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructsUint64Slice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			"no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"array no ptr",
			args{
				[2]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
		{
			"array ptr",
			args{
				[2]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]uint64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Field(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Field() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructStringSlice(t *testing.T) {
	type args struct {
		s         interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"integer no ptr",
			args{
				[]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"integer ptr",
			args{
				[]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"integer no ptr",
			args{
				[]StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"integer ptr",
			args{
				[]*StructUint16{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"string no ptr",
			args{
				[]*StructString{{"1", "1"}, {"2", "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"string ptr",
			args{
				[]*StructString{{"1", "1"}, {"2", "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"Float32 no ptr",
			args{
				[]StructFloat32{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
		{
			"Float32 ptr",
			args{
				[]*StructFloat32{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
		{
			"Float64 no ptr",
			args{
				[]StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
		{
			"Float64 ptr",
			args{
				[]*StructFloat64{{1.1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1.1", "2"},
		},
		{
			"array integer no ptr",
			args{
				[2]StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
		{
			"array integer ptr",
			args{
				[2]*StructInt8{{1, "1"}, {2, "2"}},
				"UID",
			},
			[]string{"1", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringField(tt.args.s, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	require.Panics(t, func() {
		NewStructSlice("aa")
	})

	one, two := "1", "2"
	sli := NewStructSlice([]StructMuch{{&one, errors.New("1")}, {&two, errors.New("2")}})
	t.Log(sli.Name())
	require.Panics(t, func() {
		sli.IntField("Err")
	})

	require.Panics(t, func() {
		sli.StringField("Err")
	})
	require.Panics(t, func() {
		sli.StringField("NotExist")
	})

	sli1 := NewStructSlice([]int{1, 2, 3})
	require.Panics(t, func() {
		sli1.StringField("not a struct or pointer of struct")
	})

	sli2 := NewStructSlice([]*StructMuch{nil})
	require.Panics(t, func() {
		sli2.StringField("UID")
	})
}

func TestIntSlice(t *testing.T) {
	require.Panics(t, func() {
		Int([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintSlice(t *testing.T) {
	require.Panics(t, func() {
		Uint([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Slice(t *testing.T) {
	require.Panics(t, func() {
		Int8([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int8
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int8{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int8{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint8([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint8
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint8{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint8{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint8{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Slice(t *testing.T) {
	require.Panics(t, func() {
		Int16([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int16
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int16{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int16{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint16([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint16
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint16{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint16{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint16{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Slice(t *testing.T) {
	require.Panics(t, func() {
		Int32([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int32
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int32{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int32{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint32([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint32
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint32{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint32{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint32{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Slice(t *testing.T) {
	require.Panics(t, func() {
		Int64([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []int64
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]int64{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]int64{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]int64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice(t *testing.T) {
	require.Panics(t, func() {
		Uint64([]string{"1", "2"})
	})
	tests := []struct {
		name string
		s    interface{}
		want []uint64
	}{
		{
			"int",
			[]int{1, 2, 3, 4},
			[]uint64{1, 2, 3, 4},
		},
		{
			"int",
			[]uint{1, 2, 3, 4},
			[]uint64{1, 2, 3, 4},
		},
		{
			"int",
			[]float64{1.1, 2.1, 3, 4},
			[]uint64{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSlice(t *testing.T) {
	require.Panics(t, func() {
		String([]struct{}{{}})
	})
	tests := []struct {
		name string
		s    interface{}
		want []string
	}{
		{
			"integer no ptr",
			[]int{1, 2},
			[]string{"1", "2"},
		},
		{
			"integer no ptr",
			[]uint{1, 2},
			[]string{"1", "2"},
		},
		{
			"integer no ptr",
			[]string{"1", "2"},
			[]string{"1", "2"},
		},
		{
			"Float32 no ptr",
			[]float32{1.1, 2.2},
			[]string{"1.1", "2.2"},
		},
		{
			"Float64 no ptr",
			[]float64{1.1, 2.2},
			[]string{"1.1", "2.2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringField() = %v, want %v", got, tt.want)
			}
		})
	}
}
