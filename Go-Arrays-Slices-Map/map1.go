package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// Mapas devem ser inicializados antes de serem usados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[98765432100] = "Pedro"
	aprovados[95135745682] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[95135745682])

	delete(aprovados, 95135745682)

	fmt.Println(aprovados[95135745682]) // Após a remoção, retornará uma string vazia
}
