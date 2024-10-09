package algorithms

import (
	"time"
)

func InsertionSort(arr []int) (int, int, time.Duration) {
	c := 0
	t := 0
	start := time.Now()
	for i := 0; i < len(arr); i++ {
		c++

		item := arr[i]
		j := i - 1
		for {
			c++
			if j < 0 || arr[j] < item {
				break
			}

			arr[j+1] = arr[j]
			t++
			j--
		}
		arr[j+1] = item
		t++
	}

	return c, t, time.Since(start)
}
