package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("there was an error reading the file\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	left := []int{}
	right := []int{}

	i := 0
	for scanner.Scan() {
		if i%2 == 0 {
			val, _ := strconv.Atoi(scanner.Text())
			left = append(left, val)
			i++
			continue
		}
		val, _ := strconv.Atoi(scanner.Text())
		right = append(right, val)
		i++
	}
	sort.Ints(left)
	sort.Ints(right)

	distancesSum := 0
	for i := 0; i < len(left); i++ {
		sub := (left[i] - right[i])
		if sub < 0 {
			sub *= -1
		}
		distancesSum += sub
	}
	fmt.Print(distancesSum)
}
