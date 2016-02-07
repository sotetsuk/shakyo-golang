package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCounts(s string) map[string]int {
	ret := make(map[string]int)
	for _, x := range strings.Fields(s) {
		ret[x] += 1
	}

	return ret
}

func main() {
	wc.Test(WordCounts)
}
