package gotools

// SumInt do summary of all int paras.
//
// Warning: if no para provided, zero will be returned.
// this func does not care about overflow, using at own risk.
func SumInt(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

// SumIntSlice do summary of the int slice.
//
// Warning: if empty or nil slice is provided, zero will be returned.
// this func does not care about overflow, using at own risk.
func SumIntSlice(s []int) int {
	return SumInt(s...)
}

// MaxInt guess the max int value.
func MaxInt() int {
	var v int
	v = 1
	var bits uint = 1
	for {
		v <<= 1
		if v < 0 {
			break
		}
		bits++
	}

	return (1 << bits) - 1
}
