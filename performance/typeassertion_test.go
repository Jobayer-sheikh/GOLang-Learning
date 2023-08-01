package performance

import (
	"testing"
)

type integer int64

type Incrementer interface {
	increment()
}

func (i *integer) increment() {
	*i = *i + 1
}

func BenchmarkPointer(b *testing.B) {
	i := new(integer)
	incPointer(i, b.N)
}

func BenchmarkInterface(b *testing.B) {
	i := new(integer)
	incInterface(i, b.N)
}

func BenchmarkTypeSwitch(b *testing.B) {
	i := new(integer)
	incSwitch(i, b.N)
}

func BenchmarkTypeAssertion(b *testing.B) {
	i := new(integer)
	incAssertion(i, b.N)
}

func incPointer(i *integer, n int) {
	for k := 0; k < n; k++ {
		i.increment()
	}
}

func incInterface(any Incrementer, n int) {
	for k := 0; k < n; k++ {
		any.increment()
	}
}

func incSwitch(any Incrementer, n int) {
	for k := 0; k < n; k++ {
		switch v := any.(type) {
		case *integer:
			v.increment()
		}
	}
}

func incAssertion(any Incrementer, n int) {
	for k := 0; k < n; k++ {
		if x, ok := any.(*integer); ok {
			x.increment()
		}
	}
}
