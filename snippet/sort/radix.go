package sort

import (
	"math"
)

func digitalWidth(n int) int {
	if n == 0 {
		return 1
	} else {
		var width = 0
		for n > 0 {
			width++
			n /= 10
		}
		return width
	}
}

func digitalAt(n int, k int) int {
	return (n / int(math.Pow10(k))) % 10
}

func maxDigitalWidth(nums []int) int {
	maxWidth := 0
	for i := 0; i < len(nums); i++ {
		if width := digitalWidth(nums[i]); width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth
}
func RadixSort(nums []int) {
	maxWidth := maxDigitalWidth(nums)
	help := make([]int, len(nums))
	zeroCount := make([]int, 10)
	count := make([]int, 10)
	for k := 0; k < maxWidth; k++ {
		copy(count, zeroCount)
		for _, num := range nums {
			count[digitalAt(num, k)]++
		}
		for j := 1; j < 10; j++ {
			count[j] += count[j-1]
		}
		for i := len(nums) - 1; i > -1; i-- {
			d := digitalAt(nums[i], k)
			help[count[d]-1] = nums[i]
			count[d]--
		}
		copy(nums, help)
	}
}
