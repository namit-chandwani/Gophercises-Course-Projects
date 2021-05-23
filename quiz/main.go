package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	csvFlagName  = "csv"
	defaultCSV   = "problems.csv"
	csvFlagUsage = `Name of the CSV file to be parsed`

	timeFlagName     = "limit"
	defaultTimeLimit = 30
	timeFlagUsage    = `Time limit of quiz in seconds`
)

// problem represents a (question,answer) combination from the input CSV file
type problem struct {
	question string
	answer   int
}

func main() {
	fileName := flag.String(csvFlagName, defaultCSV, csvFlagUsage)
	timeLimit := flag.Int(timeFlagName, defaultTimeLimit, timeFlagUsage)
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
	score, err := runQuiz(lines, *timeLimit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nYou have scored %d out of %d.\n", score, questionCount)
}

// runQuiz performs the quiz and returns the user score after quiz completion
func runQuiz(lines [][]string, timeLimit int) (int, error) {
	var score, input int
	inputChannel := make(chan int)

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, sum := range lines {
		answer, err := strconv.Atoi(sum[1])
		if err != nil {
			return -1, err
		}
		line := problem{question: sum[0], answer: answer}
		fmt.Printf("Problem #%d: %s = ", i+1, line.question)

		go func() {
			fmt.Scanf("%d\n", &input)
			inputChannel <- input
		}()

		select {
		case <-timer.C:
			fmt.Println("\n\nTime's up!")
			return score, nil
		case input := <-inputChannel:
			if input == line.answer {
				score++
			}
		}
	}
	return score, nil
}
