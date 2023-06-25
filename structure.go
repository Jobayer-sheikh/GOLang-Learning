package main

import "fmt"

type student struct {
	name string
	id   int
}

func (s *student) getName() string {
	return "Name: " + s.name
}

func (s *student) getPointerName() {
	s.name = "No name " + s.name
}

func main() {
	s := student{name: "Jobayer", id: 222}
	fmt.Println(s)

	ptr := &s
	ptr.id = 333

	fmt.Println(s)

	x := student{name: "hello"}
	p := &x
	fmt.Println(&p)

	fmt.Println(s.getName())
	s.getPointerName()
	fmt.Println(s.name)
}
