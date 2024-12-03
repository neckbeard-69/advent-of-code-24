package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("error while opening the file: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	println(scanner.Text())
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		stringSubs := strings.Split(line, " ")
		isSafe := true
		state := ""
		for i, sub := range stringSubs {
			if i+1 == len(stringSubs) {
				if isSafe {
					safe++
					println(line, " is safe - safe count: ", safe)
				}
				break
			}
			num1, _ := strconv.Atoi(sub)
			num2, _ := strconv.Atoi(stringSubs[i+1])

			difference := num1 - num2
			if i == 0 {
				if difference == 0 {
					println(line, " is unsafe dif is 0")
					isSafe = false
					break
				}
				if difference > 0 {
					state = "inc"
				} else {
					state = "dec"
				}

			}

			if difference > 3 || difference < -3 || difference == 0 {
				println(line, " is unsafe dif is out of bounds")
				isSafe = false
				break
			}

			if (state == "dec" && difference > 0) || (state == "inc" && difference < 0) {
				println(line, " is unsafe state")
				isSafe = false
				break
			}
		}
	}

	println(safe)
}
