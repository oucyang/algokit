package sort

import (
	"sort"
	"testing"

	"github.com/oucyang/algokit/snippet/utils"
)

func testSortInterface(t *testing.T, testTimes, maxLen, maxValue int, sorter func(p sort.Interface)) {
	for k := 0; k < testTimes; k++ {
		nums := utils.RandomIntArray(maxLen, maxValue)
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

func testSortInts(t *testing.T, testTimes, maxLen, maxValue int, sorter func([]int)) {
	for k := 0; k < testTimes; k++ {
		nums := utils.RandomIntArray(maxLen, maxValue)
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

func TestHeapSort(t *testing.T) {
	testSortInts(t, 500000, 1000, 10000, HeapSort)
}

func TestRadixSort(t *testing.T) {
	var testTimes = 50000
	var maxLen = 1000
	var maxValue = 10000

	for k := 0; k < testTimes; k++ {
		nums := utils.RandomPositiveIntArray(maxLen, maxValue)
		replica := utils.CopyIntSlice(nums)
		RadixSort(replica)
		if utils.IsSameIntSlice(nums, replica) == false {
			t.Fatalf("change nums %+v to %+v", nums, replica)
		}
		if sort.IsSorted(sort.IntSlice(replica)) == false {
			t.Fatalf("fail to sort nums %+v", nums)
		}
	}
}
