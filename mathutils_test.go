package gotools

import (
	"testing"
)

func Test_SumInt(t *testing.T) {
	if 0 != SumInt() {
		t.FailNow()
	}
	if 10 != SumInt(1, 8, 1) {
		t.FailNow()
	}
}

func Test_SumIntSlice(t *testing.T) {
	if 0 != SumIntSlice(nil) {
		t.FailNow()
	}
	if 0 != SumIntSlice([]int{}) {
		t.FailNow()
	}
	if 10 != SumIntSlice([]int{1, 8, 1}) {
		t.FailNow()
	}
}
