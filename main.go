package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Question struct {
	question string
	anwser   string
}

func main() {
	filePath := flag.String(
		"file",
		"problem.csv",
		"a path to a csv file with question and anwser",
	)

	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Total of questions: %d\n\n", len(records))

	correctScore := 0
	wrongScore := 0

	for _, q := range records {
		fmt.Printf("Question: %s\n", q[0])

		scanner.Scan()

		if scanner.Text() == q[1] {
			correctScore++
		} else {
			wrongScore++
		}
	}

	fmt.Printf("\nHow many you was correct: %d\n", correctScore)
	fmt.Printf("How many you got wrong: %d\n", wrongScore)
}
