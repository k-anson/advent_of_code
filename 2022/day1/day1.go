package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	topThreeElfCalories := [3]int{0, 0, 0}
	var currElf []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			// Empty line, finalize current elf and reset
			calorieCount := 0
			for _, calorieItem := range currElf {
				calorieCount += calorieItem
			}

			for i := 0; i < 3; i++ {
				if calorieCount > topThreeElfCalories[i] {
					temp := calorieCount
					calorieCount = topThreeElfCalories[i]
					topThreeElfCalories[i] = temp
				}
			}

			currElf = nil
		} else {
			intLine, err := strconv.Atoi(line)
			checkError(err)
			currElf = append(currElf, intLine)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Println(topThreeElfCalories)
	log.Println(topThreeElfCalories[0] + topThreeElfCalories[1] + topThreeElfCalories[2])
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
