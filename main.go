package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Question struct {
	question string
	anwser   string
}

func main() {
	file, err := os.Open("problem.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Total of questions: %d\n\n", len(records))

	score := 0

	for _, q := range records {
		fmt.Printf("Question: %s\n", q[0])

		scanner.Scan()

		if scanner.Text() == q[1] {
			fmt.Println("Nice done! You got it right.")
			score++
		} else {
			fmt.Println("Unfortunilly you wrong on that.")
		}
	}

	fmt.Printf("Your score points are: %d\n", score)
}
