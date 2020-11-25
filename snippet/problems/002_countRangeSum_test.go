package problems

import (
	"math/rand"
	"testing"
)

var slowWay = func(nums []int, lower, upper int) int {
	var ans = 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum >= lower && sum <= upper {
			ans++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				ans++
			}
		}
	}
	return ans
}

type CountRangeSumCase struct {
	Lower, Upper int
	Answer       int
	Nums         []int
}

func randomCountRangeSumCase(maxSize, maxValue int) *CountRangeSumCase {
	length := rand.Intn(maxSize)
	nums := make([]int, length)
	double := maxValue << 1
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(double) - maxValue
	}
	quadruple := double << 1
	lower := rand.Intn(quadruple) - double
	upper := rand.Intn(quadruple) - double
	if lower > upper {
		lower, upper = upper, lower
	}
	answer := slowWay(nums, lower, upper)
	return &CountRangeSumCase{
		Lower:  lower,
		Upper:  upper,
		Answer: answer,
		Nums:   nums,
	}
}

func TestCountRangeSum(t *testing.T) {
	case1 := &CountRangeSumCase{
		Lower: -2,
		Upper: 2,
		Nums:  []int{-2, 5, -1},
	}
	case2 := &CountRangeSumCase{
		Lower: -1,
		Upper: 0,
		Nums:  []int{2147483647, -2147483648, -1, 0},
	}

	var testOneCase = func(c *CountRangeSumCase) {
		ans := countRangeSum(c.Nums, c.Lower, c.Upper)
		if ans != c.Answer {
			t.Fatalf("case nums=%+v lower=%d upper=%d ans=%d right=%d", c.Nums, c.Lower, c.Upper, ans, c.Answer)
		}
	}

	testOneCase(case1)
	testOneCase(case2)

	const N = 50000
	for i := 0; i < N; i++ {
		testOneCase(randomCountRangeSumCase(1000, 1000))
	}

}
