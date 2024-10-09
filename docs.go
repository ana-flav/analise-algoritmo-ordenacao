package main

import a "main/algoritms"

var listDistributions = map[int]string{
	0: "Ordem aleatória",
	1: "Ordem crescente",
	2: "Ordem decrescente",
}

var listSizes = map[int]string{
	0: "1000",
	1: "10000",
	2: "50000",
	3: "100000",
}

var mapAlgorithms = map[string]a.SortAlg{
	"Bubble Sort":    a.BubbleSort,
	"Selection Sort": a.SelectionSort,
	"Insertion Sort": a.InsertionSort,
	"Merge Sort":     a.MergeSort,
	"Quick Sort":     a.QuickSort,
	"Heap Sort":      a.HeapSort,
}
