package gotools

import (
	"fmt"
	"reflect"
	"strconv"
)

// Str2Int convert string to int.
func Str2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

// Int2Str convert int to string.
func Int2Str(i int) string {
	return strconv.Itoa(i)
}

// ToInt convert string and any other int/float types to int.
// Warning: the result my be truncated if number overflows int.
func ToInt(v interface{}) (int, error) {
	switch v.(type) {
	case string:
		return Str2Int(v.(string))
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(v).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(v).Uint()), nil
	case float32, float64:
		return int(reflect.ValueOf(v).Float()), nil
	}

	return -1, fmt.Errorf("cannot parse value: %v", v)
}

// ToIntWithDefault tries to convert string and any other int types to int,
// return a default value provided if error occurs.
// Warning: the result my be truncated if number overflows int.
func ToIntWithDefault(v interface{}, defVal int) int {
	ret, err := ToInt(v)
	if err == nil {
		return ret
	}

	return defVal
}
