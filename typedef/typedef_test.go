package typedef

import (
	"strconv"
	"testing"
)

type StringInt string

func (s StringInt) Integer() int64 {
	i, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func TestTypedef(t *testing.T) {
	si := StringInt("1234")
	if si.Integer() != 1234 {
		t.Fatalf("got %d", si.Integer())
	}
}
