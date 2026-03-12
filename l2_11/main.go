package main

import (
	"fmt"
	"slices"
	"strings"
)

func anagram(words []string) map[string][]string {
	groups := make(map[string][]string)

	for _, w := range words {
		word := strings.ToLower(w)

		r := []rune(word)
		slices.Sort(r)

		key := string(r)

		groups[key] = append(groups[key], word)
	}

	result := make(map[string][]string)

	for _, group := range groups {
		if len(group) < 2 {
			continue
		}

		slices.Sort(group)
		result[group[0]] = group
	}

	return result
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}

	fmt.Println(anagram(words))
}
