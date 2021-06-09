package test

import (
	"errors"
	"testing"
)

func TestJSONTag(t *testing.T) {
	// t.Log(f())
	err := f()
	if err != nil {
		t.Log("err", err)
	}
}

func f() (err error) {
	// return errors.New("errors.New")
	defer func() {
		err = errors.New("defer errors.New")
		panic("panic")
	}()
	return
}
