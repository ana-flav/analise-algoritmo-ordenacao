package algorithms

import "time"

var comparacoes *int = new(int)
var trocas *int = new(int)

func merge(arr1 []int, arr2 []int) []int {
	var i, j, k int = 0, 0, 0
	lenArr3 := len(arr1) + len(arr2)
	arr3 := make([]int, 0, lenArr3)

	for {
		*comparacoes++
		if i > len(arr1) && j > len(arr2) {
			break
		}

		*comparacoes++
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			*trocas++
			k++
			i++
		} else if arr2[j] < arr1[i] {
			*comparacoes++
			arr3 = append(arr3, arr2[j])
			*trocas++
			k++
			j++
		} else {
			arr3 = append(arr3, arr1[i])
			arr3 = append(arr3, arr2[j])
			*trocas = *trocas + 2
			k = k + 2
			i++
			j++
		}

		*comparacoes++
		if i >= len(arr1) {
			arr3 = append(arr3, arr2[j:]...)
			*trocas = len(arr2[i:])
			break
		}

		*comparacoes++
		if j >= len(arr2) {
			arr3 = append(arr3, arr1[i:]...)
			*trocas = len(arr1[i:])
			break
		}
	}

	return arr3
}

func mergeSort(arr []int) []int {
	e := len(arr)
	m := e / 2

	if len(arr) == 1 {
		return arr
	}

	resultArr1 := mergeSort(arr[0:m])
	resultArr2 := mergeSort(arr[m:e])
	return merge(resultArr1, resultArr2)
}

func MergeSort(arr []int) (int, int, time.Duration) {
	start := time.Now()
	mergeSort(arr)
	return *comparacoes, *trocas, time.Since(start)
}
