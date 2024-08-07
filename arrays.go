package main

import "fmt"

type Celula struct {
	item    int
	proximo *Celula
}

func percorrerLista() {
	quinta := Celula{
		item:    5,
		proximo: nil,
	}

	quarta := Celula{
		item:    4,
		proximo: &quinta,
	}

	terceira := Celula{
		item:    3,
		proximo: &quarta,
	}

	segunda := Celula{
		item:    2,
		proximo: &terceira,
	}

	primeira := Celula{
		item:    1,
		proximo: &segunda,
	}

	atual := &primeira
	for atual != nil {
		fmt.Println(atual.item)
		atual = atual.proximo 
	}
}

func main() {
	percorrerLista()}
