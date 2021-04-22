package main

func leastBricks(wall [][]int) int {
	wm := map[int]int{}
	for _, ws := range wall {
		var tw int
		for i := 0; i < len(ws)-1; i++ {
			tw += ws[i]
			wm[tw]++
		}
	}
	var max int
	for _, t := range wm {
		if max < t {
			max = t
		}
	}
	return len(wall) - max
}
