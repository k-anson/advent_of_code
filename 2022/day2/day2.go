package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		round := scanner.Text()
		player := strings.Split(round, " ")[1]

		// Calculate win or loss points
		switch round {
		case "A X", "B Y", "C Z":
			score += 3
		case "A Y", "B Z", "C X":
			score += 6
		}

		// Calculate shape selection points
		switch player {
		case "X": // Rock
			score += 1
		case "Y": // Paper
			score += 2
		case "Z": // Scissors
			score += 3
		}
	}

	log.Println(score)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
