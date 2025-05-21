package main

import (
	csv "encoding/csv"
	"fmt"
	log "log"
	os "os"
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
	reacords, _ := reader.ReadAll()

	fmt.Println(reacords)
}
