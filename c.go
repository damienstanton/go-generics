package main

import "fmt"

// Stringer is a type constraint that requires the type argument to have
// a String method and permits the generic function to call String.
// The String method should return a string representation of the value.
type Stringer interface {
	String() string
}

type foo struct{}

func (f foo) String() string { return "I am a foo" }

func main() {
	f := foo{}
	fmt.Println(f.String())
}
