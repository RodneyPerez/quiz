package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "file", "problems.csv", "Path of quiz file")
	flag.Parse()
}

func main() {
	fmt.Println("Currently taking quiz located at this path ", fileName)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvFile := csv.NewReader(file)
	sliceOfLines, err := csvFile.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	numberOfQuestions := 0
	rightAnswers := 0

	for _, line := range sliceOfLines {
		fmt.Println("Question:", line[0])
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter answer: ")
		input, _ := reader.ReadString('\n')
		cleanedInput := strings.TrimSpace(input)
		if cleanedInput == line[1] {
			rightAnswers = rightAnswers + 1
		}
		numberOfQuestions = numberOfQuestions + 1
	}
	score := (rightAnswers / numberOfQuestions) * 100
	fmt.Println("This quiz had the following number of questions", numberOfQuestions)
	fmt.Println("You got the following right", rightAnswers)
	fmt.Println("Your Score was", score)
}
