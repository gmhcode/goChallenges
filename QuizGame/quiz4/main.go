package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	timeLimit := flag.Int("time", 30, "sets the time limit for the quiz")
	flag.Parse()
	file, _ := os.Open("problems.csv")
	problemsReader := csv.NewReader(file)

	fileContents, _ := problemsReader.ReadAll()

	problems := parseFileContents(fileContents)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	_ = timer

	newScanner := bufio.NewScanner(os.Stdin)
	_ = newScanner

	amountCorrect := 0
	answerChan := make(chan string)
	for _, problem := range problems {
		go func() {
			fmt.Println("answer this question: ", problem.question)
			newScanner.Scan()
			if newScanner.Text() == problem.answer {
				amountCorrect++
			}
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d. \n", amountCorrect, len(problems))
			return
		case <-answerChan:
		}
	}

	fmt.Println("you got ", amountCorrect, "out of 13 correct")
}

func parseFileContents(problems [][]string) []Problem {

	problemArray := make([]Problem, len(problems))

	for i, problem := range problems {
		problemArray[i].question = problem[0]
		problemArray[i].answer = problem[1]
	}

	return problemArray
}

type Problem struct {
	question string
	answer   string
}
