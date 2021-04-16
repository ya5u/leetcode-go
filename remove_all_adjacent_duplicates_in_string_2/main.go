package main

import "strings"

func removeDuplicates(s string, k int) string {
	m := make(map[string]struct{})
	for _, c := range s {
		m[string(c)] = struct{}{}
	}
	ret := s
	for {
		l := len(ret)
		for c := range m {
			ret = strings.ReplaceAll(ret, strings.Repeat(c, k), "")
		}
		if len(ret) == l {
			return ret
		}
	}
}
