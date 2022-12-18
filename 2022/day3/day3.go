package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	checkError(err)
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		firstCompartment := line[0 : len(line)/2]
		secondCompartment := line[len(line)/2:]
		itemTypes := make(map[rune]bool)

		for _, item := range firstCompartment {
			itemTypes[item] = true
		}
		for _, item := range secondCompartment {
			if itemTypes[item] {
				sum += int(getValue(item))
				break
			}
		}
	}

	log.Println(sum)
}

func getValue(char rune) rune {
	if unicode.IsLower(char) {
		return unicode.ToUpper(char) - 64
	} else {
		return char - 38
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
