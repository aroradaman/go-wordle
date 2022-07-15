package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

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

func play(tree Node, words []string) {
	var attempt, result string

	attemptHistory := make([]string, 0)

	//var attempt string
	rand.Seed(time.Now().Unix())

	word := words[rand.Intn(len(words))]
	fmt.Printf("Word: %s\n\n", word)

	fixed := make(map[rune]int)
	includes := make(map[rune][]int, 0)
	excludes := make([]rune, 0)

	for i := 0; i < 5; i++ {

		validaAttempts := make([]string, 0)

		// starting the game with "ADIEU"
		if i == 0 {
			attempt = "ADIEU"
		} else {
			pruneAndGuess(&tree, fixed, includes, excludes, make([]rune, 0), 0, len(word), &validaAttempts)
			attempt = pickOne(validaAttempts, attemptHistory)
		}

		attemptHistory = append(attemptHistory, attempt)
		result = check(word, attempt, &fixed, &includes, &excludes)

		fmt.Printf("Attempt: %s\n", attempt)
		fmt.Printf("Result: %s\n", result)

		if attempt == word {
			fmt.Printf("\nYay! completed in %d attempts\n", i+1)
			break
		}
		fmt.Printf("\nFailed to guess\n")
	}
}
