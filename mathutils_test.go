package gotools

import (
	"testing"
)

func Test_SumInt(t *testing.T) {
	if SumInt() != 0 {
		t.FailNow()
	}
	if SumInt(1, 8, 1) != 10 {
		t.FailNow()
	}
}

func Test_SumIntSlice(t *testing.T) {
	if SumIntSlice(nil) != 0 {
		t.FailNow()
	}
	if SumIntSlice([]int{}) != 0 {
		t.FailNow()
	}
	if SumIntSlice([]int{1, 8, 1}) != 10 {
		t.FailNow()
	}
}
