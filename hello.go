package main

import "fmt"

func main() {
	// Define que é uma variável e seu tipo explicitamente
	var s string
	s = "Go"

	// Define uma variável e seu tipo implicitamente utilizando o operador :=
	name := "David"

	// Define uma constante passando implicitamente o tipo
	const sobrenome = " Alves"

	// Define várias constantes ao mesmo tempo passando o tipo implicitamente
	const (
		saudacao = "Olá, tudo bem?\n"
		idade = 30
	)

	// var b bool
	// b = true

	// var i int
	// i = 42

	// var f float32
	// f = 2.93

	fmt.Printf(saudacao)
	fmt.Printf(s + " " + name + sobrenome)
	// fmt.Printf(b)
	// fmt.Printf(i)
	// fmt.Printf(f)
}