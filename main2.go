package main

import (
	"fmt"
)


func main() {
	lists := NewListsNumbers()

	// aplicando bubbleshort em listas de 1000, 10000, 50000 e 100000 em todas as variações

	fmt.Println("Lista de 1000 elementos")
	fmt.Println("Lista aleatória")
	comparacoes, trocas, duration:= bubbleSort(lists.List1000.RandomList)
	fmt.Println("Comparacoes: ", comparacoes)
	fmt.Println("Trocas: ", trocas)
	fmt.Println("Tempo: ", duration)
	

	fmt.Println("Lista ordenada")
	comparacoes, trocas, duration = bubbleSort(lists.List1000.SortedList)
	fmt.Println("Comparacoes: ", comparacoes)
	fmt.Println("Trocas: ", trocas)
	fmt.Println("Tempo: ", duration)

	fmt.Println("Lista inversa")
	comparacoes, trocas, duration= bubbleSort(lists.List1000.ReverseList)
	fmt.Println("Comparacoes: ", comparacoes)
	fmt.Println("Trocas: ", trocas)
	fmt.Println("Tempo: ", duration)
	


}
