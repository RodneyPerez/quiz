package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
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

	for _, line := range sliceOfLines {
		fmt.Println(line[0])
		fmt.Println(line[1])
	}
}
