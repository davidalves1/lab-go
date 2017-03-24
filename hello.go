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
		saudacao = "Olá, tudo bem?"
		idade = 30
	)

	// var b bool
	// b = true

	// var f float32
	// f = 2.93

	fmt.Println(saudacao)
	fmt.Println(s + " " + name + sobrenome)

	a, b := 3.0, 2.0

	fmt.Println(a + b)
	fmt.Println(a * b)
	fmt.Println(a - b)
	fmt.Println(a / b)

}