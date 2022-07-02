package main

import "unicode"

func checkFixed(word string, r rune, index int) bool {
	for i, w := range word {
		if i == index && w == r {
			return true
		}
	}
	return false
}

func checkIncludes(word string, r rune) bool {
	return !checkExcludes(word, r)
}

func checkExcludes(word string, r rune) bool {
	for _, w := range word {
		if w == r {
			return false
		}
	}
	return true
}

func check(word string, attempt string, fixed *map[rune]int, includes *map[rune][]int, excludes *[]rune) string {

	result := make([]rune, len(word))
	indices := make([]int, 0)

	// iterate over attempt for excludes and fixed first
	for i, a := range attempt {
		if checkExcludes(word, a) {
			result[i] = '_'
			*excludes = append(*excludes, a)
		} else if checkFixed(word, a, i) {
			(*fixed)[a] = i
			result[i] = a
			indices = append(indices, i)
		}
	}

	for i, a := range attempt {
		if notIn(i, indices) && checkIncludes(word, a) {
			result[i] = unicode.ToLower(a)
			_, ok := (*includes)[a]
			if ok {
				(*includes)[a] = append((*includes)[a], i)
			} else {
				(*includes)[a] = []int{i}
			}
		}
	}
	return string(result)
}
