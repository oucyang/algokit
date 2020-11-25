package sort

func MergeSort(objs []int) {
	if len(objs) < 2 {
		return
	}
	N := len(objs)
	mergeSize := 1
	for mergeSize < N {
		l := 0
		for l < N {
			m := l + mergeSize - 1
			if m >= N {
				break
			}
			r := m + mergeSize
			if r >= N {
				r = N - 1
			}
			merge(objs, l, m, r)
			l = r + 1
		}
		if mergeSize > N/2 {
			break
		}
		mergeSize <<= 1
	}
}

func MergeSortRecur(objs []int) {
	if len(objs) < 2 {
		return
	}
	mergeSortRecur(objs, 0, len(objs)-1)
}

func mergeSortRecur(objs []int, l, r int) {
	if l == r {
		return
	}
	mid := l + ((r - l) >> 1)
	mergeSortRecur(objs, l, mid)
	mergeSortRecur(objs, mid+1, r)
	merge(objs, l, mid, r)
}

func merge(objs []int, l, mid, r int) {
	var help = make([]int, r-l+1)
	var k, i, j = 0, l, mid + 1
	for i <= mid && j <= r {
		if objs[i] < objs[j] {
			help[k] = objs[i]
			k++
			i++
		} else {
			help[k] = objs[j]
			k++
			j++
		}
	}
	for i <= mid {
		help[k] = objs[i]
		k++
		i++
	}
	for j <= r {
		help[k] = objs[j]
		k++
		j++
	}
	for k := 0; k < len(help); k++ {
		objs[l+k] = help[k]
	}
}
