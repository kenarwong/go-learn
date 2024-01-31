package wordcount

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for _, v := range words {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}

	return m
}

func TestWordCount() {
	wc.Test(WordCount)
}
