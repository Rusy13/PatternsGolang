package main

import (
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		sortedWord := sortString(lowerWord)
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], lowerWord)
	}

	result := make(map[string][]string)
	for _, group := range anagramGroups {
		if len(group) > 1 {
			sort.Strings(group)
			result[group[0]] = group
		}
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)

	for key, group := range anagrams {
		println(key, ":", strings.Join(group, ", "))
	}
}
