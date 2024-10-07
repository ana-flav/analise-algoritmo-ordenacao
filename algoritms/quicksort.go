package algorithms

import (
	"fmt"
	"time"
)

var comp *int = new(int)
var tr *int = new(int)

func partition(arr *[]int, s int, e int) int {
	p := (*arr)[e]
	i := s - 1

	for j := s; j < e; j++ {
		*comp++
		if (*arr)[j] <= p {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			*tr++
		}
	}
	(*arr)[i+1], (*arr)[e] = (*arr)[e], (*arr)[i+1]
	*tr++
	return i + 1
}

func quickSort(arr *[]int, s int, e int) {
	*comp++
	if s <= e {
		p := partition(arr, s, e)

		quickSort(arr, s, p-1)
		quickSort(arr, p+1, e)
	}
}

func QuickSort(arr []int) (int, int, time.Duration) {
	start := time.Now()
	var arrPtr = &arr
	quickSort(arrPtr, 0, len(arr)-1)
	if ValidateSorting(*arrPtr) {
		fmt.Println("CERTO")
	} else {
		fmt.Println("ERRADO")
	}
	return *comp, *tr, time.Since(start)
}
