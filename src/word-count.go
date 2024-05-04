package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count_map := make(map[string]int)

	for _, word := range words {
		_, ok := count_map[word]

		if ok {
			count_map[word] = count_map[word] + 1
			continue
		}

		count_map[word] = 1
	}

	return count_map
}

func RunWordCount() {
	fmt.Println(WordCount("jon vlad vlad den pen vlad  jon"))
}
