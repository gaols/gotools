package gotools

import (
	"testing"
)

func Test_ToInt(t *testing.T) {
	if v, err := ToInt("12"); err != nil || v != 12 {
		t.FailNow()
	}
	if v, err := ToInt(int8(8)); err != nil || v != 8 {
		t.FailNow()
	}
	if v, err := ToInt(int64(8)); err != nil || v != 8 {
		t.FailNow()
	}
	if v, err := ToInt(int32(8)); err != nil || v != 8 {
		t.FailNow()
	}
	if v, err := ToInt(uint8(8)); err != nil || v != 8 {
		t.FailNow()
	}
	if v, err := ToInt(uint32(8)); err != nil || v != 8 {
		t.FailNow()
	}
	if v, err := ToInt(uint64(8)); err != nil || v != 8 {
		t.FailNow()
	}
	if v, err := ToInt(8); err != nil || v != 8 {
		t.FailNow()
	}
	if _, err := ToInt(nil); err == nil {
		t.FailNow()
	}
}
