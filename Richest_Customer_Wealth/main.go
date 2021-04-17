package main

func maximumWealth(accounts [][]int) int {
	var max int
	for _, ws := range accounts {
		var t int
		for _, w := range ws {
			t += w
		}
		if t > max {
			max = t
		}
	}
	return max
}
