package ds

import "testing"

func TestStringset_Add(t *testing.T) {
	set := NewStringset()
	set.Add("1")
	set.Add("1")
	set.Add("2")
	if set.Size() != 2 {
		t.FailNow()
	}
}

func TestStringset_Clear(t *testing.T) {
	set := NewStringset()
	set.Add("1")
	set.Add("1")
	set.Add("2")
	set.Clear()
	if set.Size() != 0 {
		t.FailNow()
	}
}

func TestStringset_Contains(t *testing.T) {
	set := NewStringset()
	set.Add("1")
	if !set.Contains("1") {
		t.FailNow()
	}
	if set.Contains("2") {
		t.FailNow()
	}
}

func TestStringset_IsEmpty(t *testing.T) {
	set := NewStringset()
	if !set.IsEmpty() {
		t.FailNow()
	}
}

func TestStringset_ContainsAll(t *testing.T) {
	set := NewStringset()
	set.Add("1")
	set.Add("2")

	if !set.ContainsAll([]string{"1", "2"}) {
		t.FailNow()
	}
}

func TestStringset_RemoveAllAll(t *testing.T) {
	set := NewStringset()
	set.Add("1")
	set.Add("2")
	set.RemoveAll([]string{"1", "2"})

	if !set.IsEmpty() {
		t.FailNow()
	}
}
