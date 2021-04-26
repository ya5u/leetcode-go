package main

func countBinarySubstrings(s string) int {
	a, p, c := 0, 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			a += min(p, c)
			p, c = c, 1
		} else {
			c++
		}
	}
	return a + min(p, c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
