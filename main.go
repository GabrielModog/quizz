package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type Question struct {
	question string
	anwser   string
}

func main() {
	filePath := flag.String(
		"file",
		"problem.csv",
		"path to a csv file with question and anwser",
	)
	timeLimit := flag.Int("limit", 30, "time limit to anwser all the questions")

	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	scanner := bufio.NewScanner(os.Stdin)

  recordsLength := len(records)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	fmt.Printf("\n=======================================\n")
	fmt.Printf(
		"\tLAZY QUIZ GAME\t\n\tby @gabrielmodog\n\tusage: -h help\n")
	fmt.Printf("=======================================\n")
	fmt.Printf("- total of questions loaded: %d\t\n", recordsLength)
	fmt.Printf("=======================================\n\n")

	correctScore := 0
	wrongScore := 0

	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\n\nTime expired, pal.\nYour score is: %d out of %d\n", correctScore, recordsLength)
    os.Exit(0)
	}()

	for _, q := range records {
		fmt.Printf("Question: %s\n", q[0])

		scanner.Scan()

		if scanner.Text() == q[1] {
			correctScore++
		} else {
			wrongScore++
		}
	}

  fmt.Printf("Your score is: %d out of %d\n", correctScore, recordsLength)
}
