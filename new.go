package main

import "fmt"


func main() {
	var nome string
	var a, b int

	fmt.Printf("Digite seu nome:") 
	fmt.Scan(&nome)
	fmt.Printf("O nome é: %s\n", nome)

	fmt.Printf("Digite o primeiro número:")
	fmt.Scan(&a)

	fmt.Printf("Digite o segundo número:")
	fmt.Scan(&b)

	fmt.Printf("A soma dos números é: %d\n", a+b)
}