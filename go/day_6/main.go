package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("answers.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalAnswers := 0
	answers := make(map[rune]struct{})

	for scanner.Scan() {
		if scanner.Text() == "" {
			totalAnswers += len(answers)
			answers = make(map[rune]struct{})
			continue
		}

		for _, answer := range scanner.Text() {
			answers[answer] = struct{}{}
		}
	}

	fmt.Println("Part One:", totalAnswers)
	return
}

func partTwo() {
	file, err := os.Open("answers.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalAnswers := 0
	answers := make(map[rune]int)
	people := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			for _, v := range answers {
				if v == people {
					totalAnswers++
				}
			}

			answers = make(map[rune]int)
			people = 0

			continue
		}

		people++

		for _, answer := range scanner.Text() {
			answers[answer]++
		}
	}

	fmt.Println("Part Two:", totalAnswers)
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
