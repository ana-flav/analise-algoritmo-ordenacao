package main

import a "main/algoritms"

var listNames = []string{
	"Ordem aleatória",
	"Ordem crescente",
	"Ordem decrescente",
}

var listSizes = []string{
	"1000",
	"10000",
	"50000",
	"100000",
}

var mapAlgorithms = map[string]a.SortAlg{
	// "Bubble Sort":    a.BubbleSort,
	// "Selection Sort": a.SelectionSort,
	// "Insertion Sort": a.InsertionSort,
	"Merge Sort":     a.MergeSort,
	// "Quick Sort":     a.QuickSort,
	// "Heap Sort":      a.HeapSort,
}
