package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa) // Exibe: <nil>

	coisa = 3
	fmt.Println(coisa) // Exibe: 3

	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2) // Exibe: Opa

	coisa2 = true
	fmt.Println(coisa2) // Exibe: true

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2) // Exibe: {Golang: Explorando a Linguagem do Google}
}
