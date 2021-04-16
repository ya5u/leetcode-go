package main

func arraySign(nums []int) int {
	x := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			x *= -1
		}
	}
	return x
}
