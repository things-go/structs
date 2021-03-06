package structs

import (
	"reflect"
)

// KeysOfMap return map key slice, need map key is string,
// if is not string, or not a map, it will panic.
func KeysOfMap(m interface{}) []string {
	rv := reflect.Indirect(reflect.ValueOf(m))
	if rv.Kind() == reflect.Invalid {
		return []string{}
	}
	if rv.Kind() != reflect.Map {
		panic("KeysOfMap: require a map")
	}

	keys := rv.MapKeys()
	ss := make([]string, 0, len(keys))
	for _, key := range keys {
		key = reflect.Indirect(key)
		if key.Kind() != reflect.String {
			panic("KeysOfMap: require string type of map key")
		}
		ss = append(ss, key.String())
	}
	return ss
}

// KeysIntOfMap return map key slice, need map key is numeric.
// (int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64).
func KeysIntOfMap(m interface{}) []int64 {
	rv := reflect.Indirect(reflect.ValueOf(m))
	if rv.Kind() == reflect.Invalid {
		return []int64{}
	}
	if rv.Kind() != reflect.Map {
		panic("KeysIntOfMap: require a map")
	}

	keys := rv.MapKeys()
	ss := make([]int64, 0, len(keys))
	for _, key := range keys {
		key = reflect.Indirect(key)
		switch key.Kind() { // nolint: exhaustive
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		default:
			panic("KeysIntOfMap: require integer type of map key")
		}
		ss = append(ss, key.Int())
	}
	return ss
}
