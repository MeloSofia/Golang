package main

import "fmt"


func digNome() {
	var nome string
	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)
	fmt.Printf("O nome é: %s\n", nome)
}

func lerNumero(prompt string) int {
	var num int
	fmt.Print(prompt)
	fmt.Scanln(&num)
	return num
}

func main() {
	digNome()

	a := lerNumero("Digite o primeiro número: ")
	b := lerNumero("Digite o segundo número: ")

	fmt.Printf("A soma dos números é: %d\n", a+b)
}