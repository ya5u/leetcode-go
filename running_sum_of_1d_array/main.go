package main

func runningSum(nums []int) []int {
	ss := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			ss[i] = num
			continue
		}
		ss[i] = ss[i-1] + num
	}
	return ss
}
