package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Question struct {
	question string
	anwser   string
}

func welcome(howManyQuestions int) {
	fmt.Printf("\n=======================================\n")
	fmt.Printf(
		"\tLAZY QUIZ GAME\t\n\tby @gabrielmodog\n\tusage: -h help\n")
	fmt.Printf("=======================================\n")
	fmt.Printf("- total of questions loaded: %d\t\n", howManyQuestions)
	fmt.Printf("=======================================\n\n")
}

func loadFile(filePathname string) [][]string {
	file, err := os.Open(filePathname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

  return records
}

func validateInput(input string) string {
  return strings.TrimSpace(strings.ToLower(input))
}

func main() {
	filePath := flag.String(
		"file",
		"problem.csv",
		"path to a csv file with question and anwser",
	)
	timeLimit := flag.Int("limit", 30, "time limit to anwser all the questions")

	flag.Parse()

	questionsData := loadFile(*filePath)
	scanner := bufio.NewScanner(os.Stdin)

  questionsLength := len(questionsData)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

  welcome(questionsLength)

	score := 0

	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("\n\nTime expired, pal.\nYour score is: %d out of %d\n", score, questionsLength)
    os.Exit(0)
	}()

	for _, q := range questionsData {
		fmt.Printf("Question: %s = ", q[0])
		scanner.Scan()

    userAnwser := validateInput(scanner.Text())
    questionAnwser := validateInput(q[1])

		if userAnwser == questionAnwser {
			score++
		}
	}

  fmt.Printf("Your score is: %d out of %d\n", score, questionsLength)
}
