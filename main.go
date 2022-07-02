package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func game(tree Node, words []string) {
	var attempt, result string

	attemptHistory := make([]string, 0)

	//var attempt string
	rand.Seed(time.Now().Unix())

	word := words[rand.Intn(len(words))]
	fmt.Printf("Word: %s\n\n", word)

	fixed := make(map[rune]int)
	includes := make(map[rune][]int, 0)
	excludes := make([]rune, 0)

	result = check(word, "CRANE", &fixed, &includes, &excludes)
	attemptHistory = append(attemptHistory, "CRANE")

	fmt.Printf("Attempt: %s\n", "CRANE")
	fmt.Printf("Result: %s\n\n", result)

	for i := 1; i < 5; i++ {
		validaAttempts := make([]string, 0)
		pruneAndGuess(&tree, fixed, includes, excludes, make([]rune, 0), 0, len(word), &validaAttempts)
		attempt = pickOne(validaAttempts, attemptHistory)
		attemptHistory = append(attemptHistory, attempt)

		result = check(word, attempt, &fixed, &includes, &excludes)
		fmt.Printf("Attempt: %s\n", attempt)
		fmt.Printf("Result: %s\n", result)
		fmt.Println()
		if attempt == word {
			fmt.Printf("Yay! completed in %d attempts\n", i+1)
			break
		}
		fmt.Printf("Failed to guess\n")
	}
}

func main() {
	var words []string

	data, _ := os.ReadFile("words.json")
	_ = json.Unmarshal(data, &words)
	//fmt.Println("words", words)

	tree := createTree(words)
	game(tree, words)

}
