package gotools

import "testing"

func TestJoinSqlInParam(t *testing.T) {
	ins := []string{"1", "2"}
	if InParam(ins) != "('1','2')" {
		t.FailNow()
	}
	ins = []string{"1"}
	if InParam(ins) != "('1')" {
		t.FailNow()
	}
}

func TestGrep(t *testing.T) {
	src := `this is a line of
magic str
magic str2
this is str
this is str2
this is str3
`
	ret := Grep(src, "magic", &GrepConfig{
		Context:                3,
		FromStart:              false,
		AlwaysIncludeFirstLine: false,
	})
	if len(ret) != 6 {
		t.FailNow()
	}
	ret = Grep(src, "magic", &GrepConfig{
		Context:                0,
		FromStart:              false,
		AlwaysIncludeFirstLine: false,
	})
	if len(ret) != 2 {
		t.FailNow()
	}
	ret = Grep(src, "magic", &GrepConfig{
		Context:                1,
		FromStart:              false,
		AlwaysIncludeFirstLine: false,
	})
	if len(ret) != 4 {
		t.FailNow()
	}
	ret = Grep(src, "magic", &GrepConfig{
		Context:                1000,
		FromStart:              false,
		AlwaysIncludeFirstLine: false,
	})
	if len(ret) != 7 {
		t.FailNow()
	}
}
