package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv file in format of question, answers")
	timeLimit := flag.Int("limit", 30, "time available to answer")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s \n", *csvFileName))

	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("something is wrong with csv file: %s", *csvFileName)
	}
	problemdata := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	counter := 0
problemloop:
	for index, problem := range problemdata {
		fmt.Printf("\nProblem #%d: %s= ", index+1, problem.question)

		answerChannel := make(chan string)
		//Go Routine Used
		go func(){
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChannel <- ans
		}()

		select {
		case <-timer.C:
			break problemloop
		case answerFromChannel := <-answerChannel:
			if answerFromChannel == problem.answer {
				counter++
			}
		}
	}

	fmt.Printf("\nScore is %d out of %d\n", counter, len(problemdata))
}

// parseLines ...
func parseLines(lines [][]string) []Problems {
	r := make([]Problems, len(lines))

	// Insert data in array of Problems
	for key, line := range lines {
		r[key] = Problems{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return r
}

// Problems ...
type Problems struct {
	question string
	answer   string
}

func exit(errorString string) {
	fmt.Printf(errorString)
	os.Exit(1)
}
