package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func generateRandomArray(maxSize int, maxValue int) []int {
	if maxValue < 0 {
		maxValue = -maxValue
	}
	double := maxValue * 2
	arrLen := rand.Intn(maxSize + 1)
	var arr = make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(double) - maxValue
	}
	return arr
}

func testSortInterface(t *testing.T, testTimes, maxSize, maxValue int, sorter func(p sort.Interface)) {
	for k := 0; k < testTimes; k++ {
		nums := generateRandomArray(maxSize, maxValue)
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

func isSameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	var m = make(map[int]int, 0)
	for _, n := range a {
		m[n] = m[n] + 1
	}
	for _, n := range b {
		m[n] = m[n] - 1
	}
	for _, c := range m {
		if c != 0 {
			return false
		}
	}
	return true
}

func testSortInts(t *testing.T, testTimes, maxSize, maxValue int, sorter func([]int)) {
	for k := 0; k < testTimes; k++ {
		nums := generateRandomArray(maxSize, maxValue)
		duplicate := make([]int, len(nums))
		copy(duplicate, nums)
		sorter(duplicate)
		if isSameSlice(nums, duplicate) == false {
			t.Fatalf("change nums %+v to %+v", nums, duplicate)
		}
		if sort.IsSorted(sort.IntSlice(duplicate)) == false {
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
