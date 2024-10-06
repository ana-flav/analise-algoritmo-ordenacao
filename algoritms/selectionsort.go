
package main

import (
    "time"
)

func selectionoSort(array []int) (int, int, time.Duration) {
	n := len(array)
	comparacoes := 0
	trocas := 0

	start := time.Now()

	for i := 0; i < n-1; i++ {
		smallest := i
		for j := i + 1; j < n; j++ {
			comparacoes++
			if array[j] < array[smallest] {
				smallest = j
			}
		array[i], array[smallest] = array[smallest], array[i]
		trocas++
		}
	}

	duration := time.Since(start)

	
	return comparacoes, trocas, duration
}

