package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "file", "problems.csv", "Path of quiz file")
	flag.Parse()
}

func main() {
	file, err := os.Open(fileName)

	if err != nil {
		exit(fmt.Sprintf("failed to open file %s", fileName))
	}

	defer file.Close()

	csvFile := csv.NewReader(file)
	sliceOfLines, err := csvFile.ReadAll()

	if err != nil {
		exit("Failed to parse the given csv file")
	}

	numberOfQuestions := 0
	rightAnswers := 0
	problems := parseLines(sliceOfLines)
	for _, problem := range problems {
		fmt.Println("Question:", problem.question)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter answer: ")
		input, _ := reader.ReadString('\n')
		cleanedInput := strings.TrimSpace(input)
		if cleanedInput == problem.answer {
			rightAnswers = rightAnswers + 1
		}
		numberOfQuestions = numberOfQuestions + 1
	}
	score := (rightAnswers / numberOfQuestions) * 100
	fmt.Println("This quiz had the following number of questions", numberOfQuestions)
	fmt.Println("You got the following right", rightAnswers)
	fmt.Println("Your Score was", score)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
