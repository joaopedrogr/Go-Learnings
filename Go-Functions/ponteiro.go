package main

import "fmt"

func inc1(n int) {
	n++ // n = n + 1
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	(*n)++ // Corrigido: precisa desreferenciar antes de incrementar
}

func main() {
	n := 1
	inc1(n) // por valor
	fmt.Println(n) // n ainda é 1, porque a função inc1 não modifica o valor original

	// revisão: & usado para obter o endereço da variável
	inc2(&n) // por referência
	fmt.Println(n) // n é 2, porque a função inc2 modifica o valor original
}
