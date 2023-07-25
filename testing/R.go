package main

import (
	"fmt"
	"log"
)

type String string

func (s *String) Size() int {
	return len(*s)
}

func (s *String) Substring(st, ed int) String {
	return String(*s)
}

func stringAndRuneCompare() {
	// string,
	s := "hello你好"

	fmt.Printf("%s, type: %T, len: %d\n", s, s, len(s))
	fmt.Printf("s[%d]: %v, type: %T\n", 0, s[0], s[0])
	li := len(s) - 1 // last index,
	fmt.Printf("s[%d]: %v, type: %T\n\n", li, s[li], s[li])

	// []rune
	rs := []rune(s)
	fmt.Printf("%v, type: %T, len: %d\n", rs, rs, len(rs))
}

func main() {

	var s string = "Jobayer"
	log.Println(s)

	var r String = "aa"
	log.Println(r)

	var char []rune = []rune(s)
	log.Println(char)

	stringAndRuneCompare()
}
