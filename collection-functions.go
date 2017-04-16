package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, p func(string) bool) bool {
	for _, v := range vs {
		if p(v) {
			return true
		}
	}
	return false
}

func All(vs []string, p func(string) bool) bool {
	for _, v := range vs {
		if !p(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, p func(string) bool) []string {
	vsp := make([]string, 0)
	for _, v := range vs {
		if p(v) {
			vsp = append(vsp, v)
		}
	}
	return vsp
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
