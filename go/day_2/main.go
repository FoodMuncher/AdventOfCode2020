package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("passwords.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	correctCount := 0

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		minmax := strings.Split(fields[0], "-")
		min, err := strconv.Atoi(minmax[0])
		check(err)
		max, err := strconv.Atoi(minmax[1])
		check(err)

		letter := rune(fields[1][0])

		letterCount := 0

		for _, v := range fields[2] {
			if letter == v {
				letterCount++
			}
		}

		if letterCount >= min && letterCount <= max {
			correctCount++
		}
	}

	fmt.Println("Part One:", correctCount)
	return
}

func partTwo() {
	file, err := os.Open("passwords.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	correctCount := 0

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		position := strings.Split(fields[0], "-")
		position1, err := strconv.Atoi(position[0])
		check(err)
		position2, err := strconv.Atoi(position[1])
		check(err)

		letter := rune(fields[1][0])
		letter1 := rune(fields[2][position1-1])
		letter2 := rune(fields[2][position2-1])

		if letter1 == letter && letter2 != letter {
			correctCount++
		} else if letter1 != letter && letter2 == letter {
			correctCount++
		}
	}

	fmt.Println("Part Two:", correctCount)
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
