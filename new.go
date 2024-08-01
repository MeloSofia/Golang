package main

import "fmt"


func main() {
	var nome string
	var a, b int

	// Solicita o nome do usuário
	fmt.Println("Digite seu nome:") 
	fmt.Scanln(&nome)
	fmt.Printf("O nome é: %s\n", nome)

	// Solicita o primeiro número
	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&a)

	// Solicita o segundo número
	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&b)

	// Calcula e imprime a soma
	fmt.Printf("A soma dos números é: %d\n", a+b)
}