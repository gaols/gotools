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

// Str2Int64 convert string to int64.
func Str2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// Int2Str convert int to string.
func Int2Str(i int) string {
	return strconv.Itoa(i)
}

// Int64ToStr convert int to string.
func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

// ToInt convert string and any other int/float types to int.
// Warning: the result my be truncated if number overflows int.
func ToInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case string:
		return Str2Int(val)
	case int:
		return val, nil
	case int8:
		return int(val), nil
	case int16:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	case uint:
		return int(val), nil
	case uint8:
		return int(val), nil
	case uint16:
		return int(val), nil
	case uint32:
		return int(val), nil
	case uint64:
		return int(val), nil
	case float32:
		return int(val), nil
	case float64:
		return int(val), nil
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
