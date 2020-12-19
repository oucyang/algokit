package problems

/*
	求数组里所有平均值不大于给定值v的子数组中最长的长度
*/

func LongestSubarrayAvgLessThan(arr []int, v int) int {
	if len(arr) < 1 {
		return 0
	}
	vals := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		vals[i] = arr[i] - v
	}

	const MaxValue = 0

	length := len(arr)
	minSum := make([]int, length)
	minSumEnd := make([]int, length)
	minSum[length-1] = vals[length-1]
	minSumEnd[length-1] = length - 1
	for i := length - 2; i > -1; i-- {
		if minSum[i+1] < 0 {
			minSum[i] = vals[i] + minSum[i+1]
			minSumEnd[i] = minSumEnd[i+1]
		} else {
			minSum[i] = vals[i]
			minSumEnd[i] = i
		}
	}
	var ans, end, sum = 0, 0, 0
	for i := 0; i < length; i++ {
		for end < length && sum+minSum[end] <= MaxValue {
			sum += minSum[end]
			end = minSumEnd[end] + 1
		}
		if ans < end-i {
			ans = end - i
		}
		if i < end {
			sum -= vals[i]
		} else {
			end = i + 1
		}
	}
	return ans
}
