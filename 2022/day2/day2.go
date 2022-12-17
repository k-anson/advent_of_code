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

	// partOne(file)
	partTwo(file)
}

func partOne(file *os.File) {
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

func partTwo(file *os.File) {
	score := 0
	scanner := bufio.NewScanner(file)
	shapes := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
	}
	shapeOrder := [3]string{"A", "B", "C"}

	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		opponent := round[0]
		opponentShapeOrderIndex := shapes[opponent]
		player := round[1]

		var playerShape string
		switch player {
		case "X": // Lose (opponents shapeOrder index + 1)
			playerShape = getPlayerShape(&shapeOrder, opponentShapeOrderIndex-1)
		case "Y": // Draw (same shape)
			score += 3
			playerShape = getPlayerShape(&shapeOrder, opponentShapeOrderIndex)
		case "Z": // Win (opponents shapeOrder index - 1)
			score += 6
			playerShape = getPlayerShape(&shapeOrder, opponentShapeOrderIndex+1)
		}
		switch playerShape {
		case "A":
			score += 1
		case "B":
			score += 2
		case "C":
			score += 3
		}
	}

	log.Println(score)
}

func getPlayerShape(shapeOrder *[3]string, index int) string {
	if index == len(shapeOrder) {
		// Start array over at index 0
		return shapeOrder[0]
	} else if index == -1 {
		// Start array at the last index
		return shapeOrder[len(shapeOrder)-1]
	} else {
		return shapeOrder[index]
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
