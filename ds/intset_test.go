package ds

import "testing"

func TestIntset_Add(t *testing.T) {
	set := NewIntset()
	set.Add(1)
	set.Add(1)
	set.Add(2)
	if set.Size() != 2 {
		t.FailNow()
	}
}

func TestIntset_Clear(t *testing.T) {
	set := NewIntset()
	set.Add(1)
	set.Add(1)
	set.Add(2)
	set.Clear()
	if set.Size() != 0 {
		t.FailNow()
	}
}

func TestIntset_Contains(t *testing.T) {
	set := NewIntset()
	set.Add(1)
	if !set.Contains(1) {
		t.FailNow()
	}
	if set.Contains(2) {
		t.FailNow()
	}
}

func TestIntset_IsEmpty(t *testing.T) {
	set := NewIntset()
	if !set.IsEmpty() {
		t.FailNow()
	}
}

func TestIntset_ContainsAll(t *testing.T) {
	set := NewIntset()
	set.Add(1)
	set.Add(2)

	if !set.ContainsAll([]int{1, 2}) {
		t.FailNow()
	}
}

func TestIntset_RemoveAllAll(t *testing.T) {
	set := NewIntset()
	set.Add(1)
	set.Add(2)
	set.RemoveAll([]int{1, 2})

	if !set.IsEmpty() {
		t.FailNow()
	}
}
