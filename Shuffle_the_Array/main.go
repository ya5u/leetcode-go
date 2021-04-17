package main

func shuffle(nums []int, n int) []int {
	s := make([]int, len(nums))
	for i := 0; i < n; i++ {
		s[2*i] = nums[i]
		s[2*i+1] = nums[n+i]
	}
	return s
}
