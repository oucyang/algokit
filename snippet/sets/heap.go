package sets

import (
	"sort"
)

func heapify(objs sort.Interface, i int) {
	length := objs.Len()
	child := (i << 1) + 1 // first child is left
	for child < length {
		if right := child + 1; right < length && objs.Less(child, right) == false {
			child = right
		}
		if objs.Less(i, child) == false {
			objs.Swap(i, child)
			i = child
		} else {
			return
		}
		child = (i << 1) + 1
	}
}

func BuildHeap(objs sort.Interface) {
	last := objs.Len() - 1
	for ; last > -1; last-- {
		heapify(objs, last)
	}
}
