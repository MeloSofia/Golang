package main

import "fmt"

var nome string

func hello() {
  fmt.Printf("Hello world\n")
}

func name() {
  nome = "Sofia"
  fmt.Printf("Olá %s\n", nome)
}

func sum(a, b int) int {
  return a + b
}

func dobro() {
    preco := []int{2, 4, 6, 8, 10}
    for _, preco := range preco {
        fmt.Printf("O dobro de %d é %d\n", preco, preco*2)
    }
}

func main() {
  hello()
  name()
  fmt.Printf("Total: %d\n", sum(10, 5))
	dobro()
}