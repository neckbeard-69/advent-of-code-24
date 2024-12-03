package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rightScanner := bufio.NewScanner(file)
	rightScanner.Split(bufio.ScanWords)

	rightSimilarities := make(map[string]int)
	i := 0
	for rightScanner.Scan() {
		if i%2 == 0 {
			i++
			continue
		}

		word := rightScanner.Text()
		rightSimilarities[word]++
		i++
	}

	file.Close()
	file, err = os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	score := 0
	i = 0
	leftScanner := bufio.NewScanner(file)
	leftScanner.Split(bufio.ScanWords)
	for leftScanner.Scan() {
		if i%2 != 0 {
			i++
			continue
		}

		word := leftScanner.Text()
		value, exists := rightSimilarities[word]
		if exists {
			wordInt, err := strconv.Atoi(word)
			if err != nil {
				log.Fatal(err)
			}
			score += (value * wordInt)
		}
		i++
	}

	println(score)
}
