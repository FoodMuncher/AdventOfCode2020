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
	file, err := os.Open("seats.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	largestSeat := 0

	for scanner.Scan() {
		lowPointer := 0
		highPointer := 127

		for i := 0; i < 7; i++ {
			if scanner.Text()[i] == 'F' {
				highPointer -= (highPointer - lowPointer + 1) / 2
			} else if scanner.Text()[i] == 'B' {
				lowPointer += (highPointer - lowPointer + 1) / 2
			}
		}

		seat := lowPointer * 8
		lowPointer = 0
		highPointer = 7

		for i := 7; i < 10; i++ {
			if scanner.Text()[i] == 'L' {
				highPointer -= (highPointer - lowPointer + 1) / 2
			} else if scanner.Text()[i] == 'R' {
				lowPointer += (highPointer - lowPointer + 1) / 2
			}
		}

		seat += lowPointer

		if seat > largestSeat {
			largestSeat = seat
		}
	}

	fmt.Println("Part One:", largestSeat)
	return
}

func partTwo() {
	file, err := os.Open("seats.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seats := make(map[int]struct{})

	for scanner.Scan() {
		lowPointer := 0
		highPointer := 127

		for i := 0; i < 7; i++ {
			if scanner.Text()[i] == 'F' {
				highPointer -= (highPointer - lowPointer + 1) / 2
			} else if scanner.Text()[i] == 'B' {
				lowPointer += (highPointer - lowPointer + 1) / 2
			}
		}

		seat := lowPointer * 8
		lowPointer = 0
		highPointer = 7

		for i := 7; i < 10; i++ {
			if scanner.Text()[i] == 'L' {
				highPointer -= (highPointer - lowPointer + 1) / 2
			} else if scanner.Text()[i] == 'R' {
				lowPointer += (highPointer - lowPointer + 1) / 2
			}
		}

		seats[seat+lowPointer] = struct{}{}
	}

	seat := 0

	for i := 1; i < 1023; i++ {
		if _, exists := seats[i]; !exists {
			if _, exists := seats[i+1]; exists {
				if _, exists := seats[i-1]; exists {
					seat = i
					break
				}
			}
		}
	}

	fmt.Println("Part One:", seat)
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
