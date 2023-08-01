package performance

import "testing"

type Title interface {
	getFullName() string
}

type Student struct {
	name string
	roll int
}

func (s Student) getFullName() string {
	return s.name
}

func (s Student) getRoll() int {
	return s.roll
}

func BenchmarkStruct(b *testing.B) {
	var s *Student = &Student{name: "Jobayer"}
	for i := 0; i < b.N; i++ {
		s.getFullName()
	}
}

func BenchmarkUsingInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s Title = &Student{name: "Jobayer"}
		s.getFullName()
	}
}

func BenchmarkUsingTypeAssertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s Title = Student{name: "Jobayer"}
		var t = s.(Student)
		t.getFullName()
	}
}
