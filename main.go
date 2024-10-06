package main

import (
	"fmt"
	a "main/algoritms"
	"reflect"
	"time"
)

var sortingAlgorithms = []a.SortAlg{a.BubbleSort, a.HeapSort, a.SelectionSort}

func executeAlgorithm(a a.SortAlg, list []int) (int, int, time.Duration) {
	return a(list)
}

func runAlgorithms(l *Lists) {
	v := reflect.ValueOf(l).Elem()

	for i := 0; i < v.NumField(); i++ {
		listas := v.Field(i)
		for j := 0; j < listas.NumField(); j++ {
			listaAny := listas.Field(j).Addr().Interface()
			lista := listaAny.(*[]int)

			for k := 0; k < len(sortingAlgorithms); k++ {
				comparacoes, trocas, duracao := executeAlgorithm(sortingAlgorithms[k], *lista)
				fmt.Printf("Tipo de lista: %v - %v, Operação: %v -- \nComparações: %v\nTrocas: %v\nDuração: %v\n\n",
					listNames[j], listSizes[i], algNames[k], comparacoes, trocas, duracao,
				)
			}
		}
	}
}

func main() {
	lists := NewListsNumbers()
	runAlgorithms(lists)
}
