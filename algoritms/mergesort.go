package main

func merge(arr1 []int, arr2 []int) []int {
	var i, j, k int = 0, 0, 0
	lenArr3 := len(arr1) + len(arr2)
	arr3 := make([]int, 0, lenArr3)

	for {
		if i > len(arr1) && j > len(arr2) {
			break
		}

		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			k++
			i++
		} else if arr2[j] < arr1[i] {
			arr3 = append(arr3, arr2[j])
			k++
			j++
		} else {
			arr3 = append(arr3, arr1[i])
			arr3 = append(arr3, arr2[j])
			k = k + 2
			i++
			j++
		}

		if i >= len(arr1) {
			arr3 = append(arr3, arr2[j:]...)
			break
		}

		if j >= len(arr2) {
			arr3 = append(arr3, arr1[i:]...)
			break
		}
	}

	return arr3
}

func MergeSort(arr []int, b int) []int {
	e := len(arr)
	m := e / 2

	if len(arr) == 1 {
		return arr
	}

	resultArr1 := MergeSort(arr[0:m], 0)
	resultArr2 := MergeSort(arr[m:e], m)
	return merge(resultArr1, resultArr2)
}
