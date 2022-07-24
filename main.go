/*
Copyright © 2022 NAME HERE <aroradaman@gmail.com>

*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strings"
	"time"

	. "github.com/daman1807/go-wordle/internal/solver"
)

var green = color.New(color.FgWhite, color.BgGreen).SprintFunc()
var yellow = color.New(color.FgWhite, color.BgYellow).SprintFunc()
var black = color.New(color.FgWhite, color.BgBlack).SprintFunc()

var cyan = color.New(color.FgCyan).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func prettyFeedback(feedback Feedback) string {
	data := make([]string, 0)

	for i := 0; i < len(feedback); i++ {
		char := feedback[i].GetVal()
		switch feedback[i].GetAnnotation() {
		case Fixed:
			data = append(data, fmt.Sprintf("%s", green(char)))
		case Included:
			data = append(data, fmt.Sprintf("%s", yellow(char)))
		case Excluded:
			data = append(data, fmt.Sprintf("%s", black(char)))
		}
	}
	return strings.Join(data, "")
}

func read(attempts int) string {
	var input string
	fmt.Printf("%s", blue(fmt.Sprintf("In [%d]: ", attempts)))
	fmt.Scanf("%s", &input)
	return strings.ToUpper(input)
}

func write(output string, exception string, attempts int) {
	if len(exception) > 0 {
		fmt.Printf("%s", red(fmt.Sprintf("Err[%d]: %s", attempts, exception)))
	} else {
		fmt.Printf("%s", cyan(fmt.Sprintf("Out[%d]: %s", attempts, output)))
	}
	fmt.Printf("\n\n")
}

func main() {

	var hints, words []string
	var answer, attempt, exception, output string
	var feedback Feedback
	var exit, completed bool

	rand.Seed(time.Now().Unix())
	data, _ := os.ReadFile("words.json")
	_ = json.Unmarshal(data, &words)

	answer = words[rand.Intn(len(words))]
	solver := NewSolver(words, answer)
	attempts := 0

	validitySolver := NewSolver(words, answer)

	fmt.Printf("Useage: [e]exit, [h]hint, defualt input will be your attempt\n\n")

	for !exit {
		attempt = read(attempts)
		output = ""
		exception = ""

		switch attempt {
		case "A":
			output = answer
		case "E":
			exit = true
			output = fmt.Sprintf("The word was: %s", answer)
		case "H":
			hints = solver.GetHints()
			if len(hints) > 0 {
				temp := make([]string, 1)
				temp[0] = "<Hints>\n"
				for i := 0; i < len(hints) && i < 5; i++ {
					temp = append(temp, fmt.Sprintf("\t%d.: %s\n", i+1, hints[i]))
				}
				output = strings.Join(temp, "")
			}
		default:

			if len(attempt) != LENGTH {
				exception = fmt.Sprintf("<expected 'e', 'h', or an attempt of length %d>", LENGTH)
			} else if !validitySolver.IsValidAttempt(attempt) {
				exception = fmt.Sprintf("<word %s not in vocabulary>", attempt)
			} else {
				feedback, completed = solver.Validate(attempt)
				output = prettyFeedback(feedback)
				attempts++
			}
		}

		write(output, exception, attempts)
		if completed {
			fmt.Printf("\nSolved in %d attempts!", attempts)
			exit = true
		}
	}
}
