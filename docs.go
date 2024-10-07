package main

import a "main/algoritms"

var listNames = []string{
	"Ordem aleatória",
	"Ordem crescente",
	"Ordem decrescente",
}

var listSizes = []string{
	"1000 itens",
	"10000 itens",
	"50000 itens",
	"100000 itens",
}

var mapAlgorithms = map[string]a.SortAlg{
	"Bubble Sort":    a.BubbleSort,
	"Selection Sort": a.SelectionSort,
	"Insertion Sort": a.InsertionSort,
	"Merge Sort":     a.MergeSort,
	"Quick Sort":     a.QuickSort,
	"Heap Sort":      a.HeapSort,
}
