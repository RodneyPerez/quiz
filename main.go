package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")

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
