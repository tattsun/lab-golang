package structs

import (
	"testing"
)

type Adder struct {
	base int
}

func (a *Adder) Add(x int) int {
	return x + a.base
}

type Muler struct {
	base int
}

func (m *Muler) Mul(x int) int {
	return x * m.base
}

type Calc struct {
	Adder
	Muler
}

func (c *Calc) Calculate(x int) int {
	return c.Add(c.Mul(x))
}

func TestStructs(t *testing.T) {
	c := Calc{
		Adder{3},
		Muler{5},
	}

	if c.Add(2) != 5 {
		t.Fatalf("got %d", c.Add(2))
	}

	if c.Mul(2) != 10 {
		t.Fatalf("got %d", c.Mul(2))
	}

	if c.Calculate(2) != 13 {
		t.Fatalf("got %d", c.Calculate(2))
	}
}
