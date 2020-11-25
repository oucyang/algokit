package sort

import "sort"

func InsertionSort(objs sort.Interface) {
	if objs == nil || objs.Len() < 2 {
		return
	}
	for i := 0; i < objs.Len(); i++ {
		for j := i - 1; j >= 0; j-- {
			if objs.Less(j+1, j) == true {
				objs.Swap(j+1, j)
			} else {
				break // [0-j] is sorted
			}
		}
	}
}
