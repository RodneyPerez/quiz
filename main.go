package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var fileName string
var timer int
var numberOfQuestions int
var rightAnswers int

func init() {
	flag.StringVar(&fileName, "file", "problems.csv", "Path of quiz file")
	flag.IntVar(&timer, "time", 30, "Time to take quiz in seconds")
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

	problems := parseLines(sliceOfLines)
	numberOfQuestions = len(problems)

	newTimer := time.NewTimer(time.Duration(timer) * time.Second)
	for _, problem := range problems {
		fmt.Println("Question:", problem.question)
		answerCh := make(chan string)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter answer: ")
			input, _ := reader.ReadString('\n')
			cleanedInput := strings.TrimSpace(input)
			answerCh <- cleanedInput
		}()
		select {
		case <-newTimer.C:
			break
		case answer := <-answerCh:
			if answer == problem.answer {
				rightAnswers++
			}
		}
	}
	finalScore(rightAnswers, numberOfQuestions)
}

func finalScore(rightAnswers int, numberOfQuestions int) {
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
			answer:   strings.TrimSpace(line[1]),
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
