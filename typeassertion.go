package main

import "fmt"

type Ducker interface {
	walk()
	quack()
}

type Duck struct{}

func (s Duck) walk()  {}
func (s Duck) quack() {}

func main() {
	var a interface{}
	a = 100
	fmt.Println(a.(int))

	a = "Hello world"
	fmt.Println(a.(string))

	var duck Ducker = Duck{}

	// Ok because Duck implement walk() behaviour
	fmt.Println(duck.(interface {
		walk()
	}))

	// Ok because Duck implement walk() and quack() behaviour
	fmt.Println(duck.(interface {
		walk()
		quack()
	}))

	// Ok because Duck implement walk() and quack() behaviour
	fmt.Println(duck.(Ducker))

	// panic: interface conversion: main.Duck is not interface { main.quack(); main.talk(); main.walk() }: missing method talk
	fmt.Println(duck.(interface {
		walk()
		quack()
		talk()
	}))
}
