package gotools

// ContainsStr returns true if arr contains val.
func ContainsStr(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

// ContainsInt returns true if arr contains val.
func ContainsInt(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

// ContainsAnyStr return true if arr contains any value in vals.
func ContainsAnyStr(arr []string, vals ...string) bool {
	for _, v := range vals {
		if ContainsStr(arr, v) {
			return true
		}
	}

	return false
}

// ContainsAnyInt return true if arr contains any value in vals.
func ContainsAnyInt(arr []int, vals ...int) bool {
	for _, v := range vals {
		if ContainsInt(arr, v) {
			return true
		}
	}

	return false
}

// ContainsAnyStrFunc return true if arr contains any value v when f(v) evalutes to true.
func ContainsAnyStrFunc(arr []string, f func(string) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}

	return false
}

// ContainsAnyIntFunc return true if arr contains any value v when f(v) evalutes to true.
func ContainsAnyIntFunc(arr []int, f func(int) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}

	return false
}

func Contains[T comparable](arr []T, val T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func ContainsFunc[T comparable](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}		
	}	
	return false
}
