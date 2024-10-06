package main

import "fmt"

func main() {
	lista := []int{1, 2, 5, 8, 9, 10, 2, 3, 5, 6, 7, 4}
	listaOrdenada := MergeSort(lista, 0)
	fmt.Printf("%v", listaOrdenada)
}
