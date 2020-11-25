package sort

import "sort"

func SelectionSort(objs sort.Interface) {
	if objs == nil || objs.Len() < 2 {
		return
	}
	for i := 0; i < objs.Len(); i++ {
		var minIndex = i
		for j := i + 1; j < objs.Len(); j++ {
			if objs.Less(j, minIndex) == true {
				minIndex = j
			}
		}
		objs.Swap(minIndex, i)
	}
}
