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
