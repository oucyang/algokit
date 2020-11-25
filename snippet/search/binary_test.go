package search

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

func generateRandomNumsSort(maxSize, maxValue int) []int {
	arrLen := rand.Intn(maxSize + 1)
	arr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Intn(maxValue + 1)
	}
	sort.Ints(arr)
	return arr
}

func TestBinarySearch(t *testing.T) {
	const N = 500000
	for k := 0; k < N; k++ {
		nums := generateRandomNumsSort(100, 1000)
		for i := 0; i < len(nums); i++ {
			target := nums[i]
			index := BinarySearch(nums, target)
			if nums[index] != target {
				t.Fatalf("fail to search target %d in %+v", target, nums)
			}
		}
	}
}

func mostLeftMoreThan(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return -1
}

func TestBinarySearchNearLeft(t *testing.T) {
	const N = 500

	var testOnce = func(nums []int, target int) {
		result := BinarySearchNearLeft(nums, target)
		rightResult := mostLeftMoreThan(nums, target)
		if result != rightResult {
			t.Fatalf("result=%d right=%d fail to search most left %d in %+v", result, rightResult, target, nums)
		}
	}

	for k := 0; k < N; k++ {
		nums := generateRandomNumsSort(100, 1000)
		before, after := math.MinInt64, math.MaxInt64
		testOnce(nums, before)
		testOnce(nums, after)
		for i := 0; i < len(nums); i++ {
			target := nums[i]
			testOnce(nums, target)
			testOnce(nums, target+1)
		}
	}
}

func mostRightLessThan(nums []int, target int) int {
	for i := len(nums) - 1; i > -1; i-- {
		if nums[i] <= target {
			return i
		}
	}
	return -1
}

func TestBinarySearchNearRight(t *testing.T) {
	const N = 500

	var testOnce = func(nums []int, target int) {
		result := BinarySearchRight(nums, target)
		rightResult := mostRightLessThan(nums, target)
		if result != rightResult {
			t.Fatalf("result=%d right=%d fail to search most left %d in %+v", result, rightResult, target, nums)
		}
	}

	for k := 0; k < N; k++ {
		nums := generateRandomNumsSort(100, 1000)
		before, after := math.MinInt64, math.MaxInt64
		testOnce(nums, before)
		testOnce(nums, after)
		for i := 0; i < len(nums); i++ {
			target := nums[i]
			testOnce(nums, target)
			testOnce(nums, target-1)
		}
	}
}

func TestHalfSum(t *testing.T) {
	var halfSum = func(a, b uint64) uint64 {
		if a > b {
			return b + (a-b)/2
		}
		return a + (b-a)/2
	}

	var testOnce = func(a, b uint64) {
		result := HalfSum(a, b)
		right := halfSum(a, b)
		if result != right {
			t.Fatalf("half of sum a=%d b=%d result=%d right=%d", a, b, result, right)
		}
	}

	testOnce(math.MaxInt64, math.MaxInt64)
	const N = 10000000
	for k := 0; k < N; k++ {
		a, b := rand.Uint64(), rand.Uint64()
		testOnce(a, b)
	}
}
