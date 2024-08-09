package main

import "fmt"

type Celula struct {
    item    int
    proximo *Celula
}
// Inserir um item na lista
func inserir(raiz **Celula, novoItem int) {
    novaCelula := &Celula{
        item:    novoItem,
        proximo: nil,
    }

    if *raiz == nil || novoItem < (*raiz).item {
        novaCelula.proximo = *raiz
        *raiz = novaCelula
        return
    }

    celulaAtual := *raiz

    for celulaAtual.proximo != nil && celulaAtual.proximo.item < novoItem {
        celulaAtual = celulaAtual.proximo
    }

    novaCelula.proximo = celulaAtual.proximo
    celulaAtual.proximo = novaCelula
}

func imprimirCelulas(celula *Celula) {
    celulaAtual := celula

    for celulaAtual != nil {
        fmt.Printf("Celula(%d) -> ", celulaAtual.item)
        celulaAtual = celulaAtual.proximo
    }

    fmt.Println("nil")
}

// Buscar um item na lista
func buscar(raiz *Celula, item int) *Celula {
    celulaAtual := raiz

    for celulaAtual != nil {
        if celulaAtual.item == item {
            return celulaAtual
        }
        celulaAtual = celulaAtual.proximo
    }
    return nil
}

// Remover um item da lista
func remover(raiz **Celula, itemParaRemover int) {
    if *raiz == nil {
        return
    }
    if (*raiz).item == itemParaRemover {
        *raiz = (*raiz).proximo
        return
    }

    anterior := *raiz
    atual := (*raiz).proximo
    for atual != nil {
        if atual.item == itemParaRemover {
            anterior.proximo = atual.proximo
            return
        }
        anterior = atual
        atual = atual.proximo
    }
}

func main() {
    var raiz *Celula = nil

    imprimirCelulas(raiz)

    inserir(&raiz, 6)
    imprimirCelulas(raiz)

    inserir(&raiz, 10)
    imprimirCelulas(raiz)

    inserir(&raiz, 11)
    imprimirCelulas(raiz)

    inserir(&raiz, 12)
    imprimirCelulas(raiz)

    inserir(&raiz, 13)
    imprimirCelulas(raiz)

    inserir(&raiz, 2)
    imprimirCelulas(raiz)

		inserir(&raiz, 5)
    imprimirCelulas(raiz)

		inserir(&raiz, 16)
    imprimirCelulas(raiz)

		itemParaBuscar := 2

    resultado := buscar(raiz, itemParaBuscar)
		if resultado != nil{
			fmt.Println("Item existente:",itemParaBuscar)
		}else{
			fmt.Println("Item inesxistente")
		}

		remover(&raiz, 11)
    fmt.Println("Lista após remover o item:")
    imprimirCelulas(raiz)
}

