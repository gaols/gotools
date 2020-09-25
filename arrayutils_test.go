package gotools

import "testing"

func Test_ContainsStr(t *testing.T) {
	if !ContainsStr([]string{"hello", "hello1"}, "hello") {
		t.FailNow()
	}
	if ContainsStr([]string{"hello", "hello1"}, "hello2") {
		t.FailNow()
	}
	if ContainsStr(nil, "hello2") {
		t.FailNow()
	}
}

func Test_ContainsInt(t *testing.T) {
	if !ContainsInt([]int{1, 2}, 1) {
		t.FailNow()
	}
	if ContainsInt([]int{1, 2}, 3) {
		t.FailNow()
	}
	if ContainsInt(nil, 1) {
		t.FailNow()
	}
}

func Test_ContainsAnyStr(t *testing.T) {
	if !ContainsAnyStr([]string{"hello", "hello1"}, "hello") {
		t.FailNow()
	}
	if ContainsAnyStr([]string{"hello", "hello1"}, "hello2") {
		t.FailNow()
	}
	if ContainsAnyStr(nil, "hello2") {
		t.FailNow()
	}
	if !ContainsAnyStr([]string{"hello2"}, "hello", "hello2") {
		t.FailNow()
	}
	if ContainsAnyStr(nil, "") {
		t.FailNow()
	}
}

func Test_ContainsAnyInt(t *testing.T) {
	if !ContainsAnyInt([]int{1, 2}, 1) {
		t.FailNow()
	}
	if ContainsAnyInt([]int{1, 2}, 0) {
		t.FailNow()
	}
	if ContainsAnyInt(nil, 1) {
		t.FailNow()
	}
	if !ContainsAnyInt([]int{0}, 0, 2) {
		t.FailNow()
	}
}
