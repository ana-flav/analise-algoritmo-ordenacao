package algorithms

import (
	"time"
)

func heapify(arr []int, n int, i int, comparacoes *int, trocas *int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		*comparacoes++
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		*comparacoes++
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		*trocas++

		heapify(arr, n, largest, comparacoes, trocas)
	}
}

func HeapSort(arr []int) (int, int, time.Duration) {
	n := len(arr)
	comparacoes := 0
	trocas := 0

	start := time.Now()

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, &comparacoes, &trocas)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		trocas++

		heapify(arr, i, 0, &comparacoes, &trocas)
	}

	duration := time.Since(start)
	return comparacoes, trocas, duration
}
