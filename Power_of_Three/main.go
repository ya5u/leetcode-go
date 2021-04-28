package main

import (
	"math"
)

func isPowerOfThree(n int) bool {
	fn := float64(n)
	i := 0
	for {
		fi := float64(i)
		p := math.Pow(3.0, fi)
		if p == fn {
			return true
		} else if p > fn {
			return false
		}
		i++
	}
}
