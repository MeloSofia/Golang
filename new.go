package main

import (
	"fmt"
)

func soma() {
	for {
		var escolha string
		var a, b float32

		fmt.Println("\nEscolha uma forma: \n• Triângulo \n• Retângulo \n• Círculo \n• Sair")
		fmt.Scan(&escolha)

		switch escolha {
		case "Triângulo":
			fmt.Printf("Digite a base: ")
			fmt.Scan(&a)
			fmt.Printf("Digite a altura: ")
			fmt.Scan(&b)
			areaT := (a * b) / 2
			fmt.Printf("A área do triângulo é: %.2f\n", areaT)

		case "Círculo":
			fmt.Printf("Digite o raio: ")
			fmt.Scan(&a)
			var pi float32 = 3.14
			areaC := (a * a) * pi
			fmt.Printf("A área do círculo é: %.2f\n", areaC)

		case "Retângulo":
			fmt.Printf("Digite a base: ")
			fmt.Scan(&a)
			fmt.Printf("Digite a altura: ")
			fmt.Scan(&b)
			areaR := a * b
			fmt.Printf("A área do retângulo é: %.2f\n", areaR)

		case "Sair":
			fmt.Println("Finalizado")
			return

		default:
			fmt.Println("Escolha inválida, tente novamente.")
		}
	}
}

func main() {
	soma()
}
// func main() {
// 	var nome string
// 	var a, b int

// 	fmt.Printf("Digite seu nome:") 
// 	fmt.Scan(&nome)
// 	fmt.Printf("O nome é: %s\n", nome)

// 	fmt.Printf("Digite o primeiro número:")
// 	fmt.Scan(&a)

// 	fmt.Printf("Digite o segundo número:")
// 	fmt.Scan(&b)

// 	fmt.Printf("A soma dos números é: %d\n", a+b)
// }
