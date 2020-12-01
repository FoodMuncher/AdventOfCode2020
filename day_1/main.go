package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("expense_report.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	expenses := make(map[int]struct{})

	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		check(err)

		if _, exists := expenses[2020-expense]; exists {
			fmt.Println("Part One:", (2020-expense)*expense)
			return
		}

		expenses[expense] = struct{}{}
	}
}

func partTwo() {
	file, err := os.Open("expense_report.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	expenses := make(map[int]struct{})

	for scanner.Scan() {
		expense1, err := strconv.Atoi(scanner.Text())
		check(err)

		for expense2 := range expenses {
			if _, exists := expenses[2020-expense1-expense2]; exists {
				fmt.Println("Part Two:", (2020-expense1-expense2)*expense1*expense2)
				return
			}
		}

		expenses[expense1] = struct{}{}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
