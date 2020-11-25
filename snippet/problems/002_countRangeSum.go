package problems

// https://leetcode.com/problems/count-of-range-sum/

func countRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, len(nums)+1)
	for i, num := range nums {
		sums[i+1] = sums[i] + num
	}
	ans := process(sums, 0, len(sums)-1, lower, upper)
	return ans
}

func process(sums []int, l, r, lower, upper int) int {
	if r == l {
		return 0
	}
	m := l + ((r - l) >> 1)
	ans := process(sums, l, m, lower, upper) +
		process(sums, m+1, r, lower, upper) +
		merge(sums, l, m, r, lower, upper)
	return ans
}

func merge(sums []int, l, m, r, lower, upper int) int {
	var ans = 0
	var help = make([]int, (r - l + 1))
	var k, i, j = 0, l, m + 1
	var jLower, jUpper = j, j
	for i <= m {
		for jLower <= r && sums[jLower]-sums[i] < lower {
			jLower++
		}
		for jUpper <= r && sums[jUpper]-sums[i] <= upper {
			jUpper++
		}
		ans += jUpper - jLower
		for j <= r && sums[j] < sums[i] {
			help[k] = sums[j]
			j++
			k++
		}
		help[k] = sums[i]
		i++
		k++
	}
	for j <= r {
		help[k] = sums[j]
		k++
		j++
	}
	for k = 0; k < len(help); k++ {
		sums[l+k] = help[k]
	}
	return ans
}
