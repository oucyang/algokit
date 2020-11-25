package sets

import (
	"github.com/oucyang/algokit/snippet/utils"
	"sort"
	"testing"
)

func isHeap(objs sort.Interface) (int, int, bool) {
	length := objs.Len()
	for i := 0; i < length; i++ {
		left := (i << 1) + 1
		if left >= length {
			return -1, -1, true
		}
		if objs.Less(left, i) == true { // `less(i, j) == true` is `i < j`, not `i <= j`
			return i, left, false
		}
		right := left + 1
		if right >= length {
			return -1, -1, true
		}
		if objs.Less(right, i) == true {
			return i, right, false
		}
	}
	return -1, -1, true
}

func TestBuildHeap(t *testing.T) {
	const N = 50000
	for i := 0; i < N; i++ {
		nums := utils.RandomIntArray(1000, 10000)
		BuildHeap(sort.IntSlice(nums))
		if parent, child, ok := isHeap(sort.IntSlice(nums)); !ok {
			t.Fatalf("nums=%+v is not heap, parent=%d, child=%d", nums, parent, child)
		}
	}
}
