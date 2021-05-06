package main

func jump(nums []int) int {
	l := len(nums)
	var a, c int
	if l == 1 {
		return a
	}
	for {
		a++
		n := nums[c]
		if c+n >= l-1 {
			return a
		}
		var max, mi int
		for i := n; i > 0; i-- {
			if max == 0 || max < i+nums[c+i] {
				max = i + nums[c+i]
				mi = i
			}
		}
		c += mi
	}
}
