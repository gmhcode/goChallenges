package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	timeLimit := flag.Int("limit", 30, "Test time limit")
	fileContent := getFileContent("problems.csv")
	flag.Parse()
	problems := parseProblems(fileContent)
	correct := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)

		answerChan := make(chan string)

		go func() {
			var answer string
			fmt.Scan(&answer)
			if answer == problem.answer {
				correct++
			}
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You got %d/%d correct \n ", correct, len(problems))
			return
			//could also do case answer := <- answerChan
		case <-answerChan:

		}
	}
	fmt.Printf("You got %d/%d correct \n ", correct, len(problems))

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func getFileContent(fileName string) [][]string {

	file, err := os.Open("problems.csv")
	if err != nil {
		exit("Could not open file")
	}

	fileReader := csv.NewReader(file)

	fileContent, err := fileReader.ReadAll()
	return fileContent
}

func readAsText() {
	dat, _ := ioutil.ReadFile("problems.csv")
	fmt.Println("dat ", string(dat))
}

func parseProblems(problems [][]string) []problem {

	returningProblems := make([]problem, len(problems))

	for i := range returningProblems {
		returningProblems[i].answer = strings.TrimSpace(problems[i][1])
		returningProblems[i].question = problems[i][0]
	}

	return returningProblems
}

type problem struct {
	question string
	answer   string
}
