package main

import (
	"fmt"
	a "main/algoritms"
	"reflect"
	"time"
)

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

			for k, v := range mapAlgorithms {
				comparacoes, trocas, duracao := executeAlgorithm(v, *lista)
				fmt.Printf("Tipo de lista: %v - %v, Algoritmo: %v -- \nComparações: %v\nTrocas: %v\nDuração: %v\n\n",
					listNames[j], listSizes[i], k, comparacoes, trocas, duracao,
				)
			}
		}
	}
}

func main() {
	lists := NewListsNumbers()
	runAlgorithms(lists)
}
