package main

import "fmt"

func main() {
	// Define que é uma variável e seu tipo explicitamente
	var s string
	s = "Go"

	// Define uma variável e seu tipo implicitamente utilizando o operador :=
	name := "David"

	// var b bool
	// b = true

	// var i int
	// i = 42

	// var f float32
	// f = 2.93

	fmt.Printf(s + " " + name)
	// fmt.Printf(b)
	// fmt.Printf(i)
	// fmt.Printf(f)
}