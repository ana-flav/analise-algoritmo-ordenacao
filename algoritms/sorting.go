package algorithms

import (
	"fmt"
	"time"
)

type SortAlg func([]int) (int, int, time.Duration)

func ValidateSorting(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			str := fmt.Sprintf("%v -- %v %v, position %v %v", arr, arr[i], arr[i+1], i, i+1)
			panic(str)
		}
	}
	return true
}
