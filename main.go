package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
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

func parseRecords(records [][]string, shuffle bool) []Problem {
	problems := make([]Problem, len(records))
	for i, p := range records {
		problems[i] = Problem{
			question: p[0],
			anwser:   validateInput(p[1]),
		}
	}

	if shuffle {
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}

	return problems
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
	shuffle := flag.Bool("shuffle", false, "randomly shuffle questions")

	flag.Parse()

	data := loadFile(*filePath)
	problems := parseRecords(data, *shuffle)

	scanner := bufio.NewScanner(os.Stdin)

	problemsAmount := len(problems)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	welcome(problemsAmount)

	score := 0

	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf(
			"\n\nTime expired, pal.\nYour score is: %d out of %d\n",
			score,
			problemsAmount,
		)
		os.Exit(0)
	}()

	for _, p := range problems {
		fmt.Printf("Question: %s = ", p.question)
		scanner.Scan()

		userAnwser := validateInput(scanner.Text())

		if userAnwser == p.anwser {
			score++
		}
	}

	fmt.Printf("Your score is: %d out of %d\n", score, problemsAmount)
}
