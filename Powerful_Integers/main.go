package main

func powerfulIntegers(x int, y int, bound int) []int {
	if x == 1 && y == 1 {
		if bound < 2 {
			return []int{}
		} else {
			return []int{2}
		}
	}
	m := map[int]struct{}{}
	xx, yy := 1, 1
	if xx+yy > bound {
		return []int{}
	}
	m[xx+yy] = struct{}{}
	if x != 1 && y != 1 {
		for {
			yy *= y
			if xx+yy > bound {
				break
			}
			m[xx+yy] = struct{}{}
		}
		for {
			yy = 1
			xx *= x
			if xx+yy > bound {
				break
			}
			m[xx+yy] = struct{}{}
			for {
				yy *= y
				if xx+yy > bound {
					break
				}
				m[xx+yy] = struct{}{}
			}
		}
	} else {
		var z int
		if x != 1 {
			z = x
		} else if y != 1 {
			z = y
		}
		zz := 1
		for {
			zz *= z
			if zz+1 > bound {
				break
			}
			m[zz+1] = struct{}{}
		}
	}
	a := make([]int, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}
