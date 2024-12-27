package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(strings.Fields(s)); i++ {
		m[strings.Fields(s)[i]] += 1;
	}
	return m;
}

func main() {
	wc.Test(WordCount)
}
