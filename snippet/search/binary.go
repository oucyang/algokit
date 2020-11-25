package search

func BinarySearch(sortedNums []int, target int) int {
	if len(sortedNums) < 1 {
		return -1
	}
	var L, R, M = 0, len(sortedNums) - 1, 0
	for L < R {
		M = L + ((R - L) >> 1)
		if sortedNums[M] == target {
			return M
		} else if sortedNums[M] > target {
			R = M - 1
		} else {
			L = M + 1
		}
	}
	if sortedNums[L] == target {
		return L
	}
	return -1
}

// the most left value more than target
func BinarySearchNearLeft(nums []int, target int) int {
	var L, R, M = 0, len(nums) - 1, 0
	var index = -1
	for L <= R {
		M = L + ((R - L) >> 1)
		if nums[M] >= target {
			index = M
			R = M - 1
		} else {
			L = M + 1
		}
	}
	return index
}

// the most right value less than target
func BinarySearchRight(nums []int, target int) int {
	var L, R, M = 0, len(nums) - 1, 0
	var index = -1
	for L <= R {
		M = L + ((R - L) >> 1)
		if nums[M] <= target {
			index = M
			L = M + 1
		} else {
			R = M - 1
		}
	}
	return index
}

func HalfSum(a, b uint64) uint64 {
	return a&b + ((a ^ b) >> 1)
}
