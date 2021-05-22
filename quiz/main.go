package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	csvFlagName  = "csv"
	defaultCSV   = "problems.csv"
	csvFlagUsage = `Name of the CSV file to be parsed`
)

// problem represents a (question,answer) combination from the input CSV file
type problem struct {
	question string
	answer   int
}

func main() {
	fileName := flag.String(csvFlagName, defaultCSV, csvFlagUsage)
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	questionCount := len(lines)
	score, err := runQuiz(lines)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You scored %d out of %d.\n", score, questionCount)
}

// runQuiz performs the quiz and returns the user score after quiz completion
func runQuiz(lines [][]string) (int, error) {
	var score, input int

	for i, sum := range lines {
		answer, err := strconv.Atoi(sum[1])
		if err != nil {
			return -1, err
		}
		line := problem{question: sum[0], answer: answer}
		fmt.Printf("Problem #%d: %s = ", i+1, line.question)
		fmt.Scanf("%d\n", &input)
		if input == line.answer {
			score++
		}
	}
	return score, nil
}
