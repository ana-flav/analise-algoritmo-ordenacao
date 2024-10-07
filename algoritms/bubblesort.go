package algorithms

import (
	"fmt"
	"time"
)

func BubbleSort(array []int) (int, int, time.Duration) {
	n := len(array)
	comparacoes := 0
	trocas := 0

	start := time.Now()

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			comparacoes++
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				trocas++
			}
		}
	}

	if ValidateSorting(array) {
		fmt.Println("CERTO")
	} else {
		fmt.Println("ERRADO")
	}

	duration := time.Since(start)

	return comparacoes, trocas, duration
}
