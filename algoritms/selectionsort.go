package algorithms

import (
	"fmt"
	"time"
)

func SelectionSort(array []int) (int, int, time.Duration) {
	n := len(array)
	comparacoes := 0
	trocas := 0

	start := time.Now()

	// {5, 6, 9, 3, 1, 10, 8, 2, 7, 4}

	for i := 0; i < n-1; i++ {
		smallest := i
		for j := i + 1; j < n; j++ {
			comparacoes++
			if array[j] < array[smallest] {
				smallest = j
			}
		}
		array[i], array[smallest] = array[smallest], array[i]
		trocas++
	}

	if ValidateSorting(array) {
		fmt.Println("CERTO")
	} else {
		fmt.Println("ERRADO")
	}



	duration := time.Since(start)

	return comparacoes, trocas, duration
}
