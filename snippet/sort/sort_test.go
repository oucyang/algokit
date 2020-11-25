package sort

import (
	"sort"
	"testing"

	"github.com/oucyang/algokit/snippet/utils"
)

func testSortInterface(t *testing.T, testTimes, maxSize, maxValue int, sorter func(p sort.Interface)) {
	for k := 0; k < testTimes; k++ {
		nums := utils.RandomIntArray(maxSize, maxValue)
		sorter(sort.IntSlice(nums))
		if sort.IsSorted(sort.IntSlice(nums)) == false {
			t.Fatalf("fail to sort nums %+v", nums)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	testSortInterface(t, 50000, 1000, 10000, BubbleSort)
}

func TestSelectionSort(t *testing.T) {
	testSortInterface(t, 50000, 1000, 10000, SelectionSort)
}

func TestInsertionSort(t *testing.T) {
	testSortInterface(t, 50000, 1000, 10000, InsertionSort)
}

func testSortInts(t *testing.T, testTimes, maxSize, maxValue int, sorter func([]int)) {
	for k := 0; k < testTimes; k++ {
		nums := utils.RandomIntArray(maxSize, maxValue)
		replica := utils.CopyIntSlice(nums)
		sorter(replica)
		if utils.IsSameIntSlice(nums, replica) == false {
			t.Fatalf("change nums %+v to %+v", nums, replica)
		}
		if sort.IsSorted(sort.IntSlice(replica)) == false {
			t.Fatalf("fail to sort nums %+v", nums)
		}
	}
}

func TestMergeSortRecur(t *testing.T) {
	testSortInts(t, 500000, 1000, 10000, MergeSortRecur)
}

func TestMergeSort(t *testing.T) {
	testSortInts(t, 500000, 1000, 10000, MergeSort)
}

func TestQuickSortV1(t *testing.T) {
	testSortInterface(t, 50000, 1000, 10000, QuickSortV1)
}

func TestQuickSortV2(t *testing.T) {
	testSortInts(t, 500000, 1000, 10000, QuickSortV2)
}

func TestQuickSortV3(t *testing.T) {
	testSortInts(t, 500000, 1000, 10000, QuickSortV3)
}
