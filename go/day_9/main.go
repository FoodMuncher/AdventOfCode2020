package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type bags []string

func main() {
	number := partOne()
	partTwo(number)
}

func partOne() int {
	file, err := os.Open("xmas.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		check(err)
		numbers = append(numbers, number)
	}

	for i := 25; i < len(numbers); i++ {
		valid := false
		number := numbers[i]

	loop:
		for j := i - 25; j < i; j++ {
			for z := j + 1; z < i; z++ {
				if numbers[j]+numbers[z] == number {
					valid = true
					break loop
				}
			}
		}

		if !valid {
			fmt.Println("Part One:", number)
			return number
		}
	}

	return -1
}

func partTwo(prevNumber int) {
	file, err := os.Open("xmas.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		check(err)
		numbers = append(numbers, number)
	}

	for size := 2; size < len(numbers); size++ {
		for startIndex := 0; startIndex < len(numbers)-size; startIndex++ {
			sum := 0

			for i := startIndex; i < startIndex+size; i++ {
				sum += numbers[i]
			}

			if sum == prevNumber {
				max := numbers[startIndex]
				min := numbers[startIndex]

				for i := startIndex; i < startIndex+size; i++ {
					if numbers[i] > max {
						max = numbers[i]
					}

					if numbers[i] < min {
						min = numbers[i]
					}
				}

				fmt.Println("Part Two:", max+min)
				return
			}
		}

	}

	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
