package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "The file name for the quiz")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz")

	file, _ := os.Open(*csvFilename)

	problems := csv.NewReader(file)

	fileContents, _ := problems.ReadAll()

	problemArray := parseProblems(fileContents)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	numberCorrect := 0

	for _, p := range problemArray {

		answerChan := make(chan string)
		var answer string
		go func() {
			fmt.Println("What is the answer to this problem: ", p.question)
			fmt.Scan(&answer)
			if answer == p.answer {
				numberCorrect++
			}
			answerChan <- answer

		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d. \n", numberCorrect, len(problemArray))
			return
		case <-answerChan:
		}
	}
	fmt.Printf("You scored %d out of %d. \n", numberCorrect, len(problemArray))
}

func parseProblems(problems [][]string) []problem {

	problemArray := make([]problem, len(problems))
	for i, p := range problems {
		problemArray[i].question = p[0]
		problemArray[i].answer = p[1]
	}

	return problemArray

type problem struct {
	question string
	answer   string
}
