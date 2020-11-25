package sort

import "sort"

func BubbleSort(objs sort.Interface) {
	if objs == nil || objs.Len() < 2 {
		return
	}
	for e := objs.Len() - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if objs.Less(i+1, i) == true {
				objs.Swap(i+1, i)
			}
		}
	}
}
