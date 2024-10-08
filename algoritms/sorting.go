package algorithms

import (
	"time"
)

type SortAlg func([]int) (int, int, time.Duration)

func ValidateSorting(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
