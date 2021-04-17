package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int
	for _, c := range candies {
		if c > max {
			max = c
		}
	}
	s := make([]bool, len(candies))
	for i, c := range candies {
		s[i] = c+extraCandies >= max
	}
	return s
}
