package main

import "fmt"


func main() {
	var nome string
	var a, b int

	fmt.Println("Digite seu nome:") 
	fmt.Scanln(&nome)
	fmt.Printf("O nome é: %s\n", nome)

	fmt.Println("Digite o primeiro número:")
	fmt.Scanln(&a)

	fmt.Println("Digite o segundo número:")
	fmt.Scanln(&b)

	fmt.Printf("A soma dos números é: %d\n", a+b)
}