package filetools

import "testing"

func TestPathExists(t *testing.T) {
	v, e := PathExists("filetools.go")
	if e != nil {
		t.Error(e)
	}
	if v == false {
		t.Fail()
	}

	v, e = PathExists("filetools1.go")
	if e != nil {
		t.Error(e)
	}
	if v {
		t.Fail()
	}
}
