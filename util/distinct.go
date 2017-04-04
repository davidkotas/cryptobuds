package util

import (
	"sort"
)

func Distinct(s []int) []int {
	h := map[int]int{}

	for _, str := range s {
		h[str]++
	}

	d := []int{}

	for k, _ := range h {
		d = append(d, k)
	}

	slice := sort.IntSlice(d)
	sort.Sort(slice)

	return d
}
