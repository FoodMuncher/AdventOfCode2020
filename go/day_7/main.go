package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bags []string

func main() {
	partOne()
	partTwo()
}

func partOne() {
	correctBags := make(map[string]struct{})
	numberOfBags := -1

	for len(correctBags) != numberOfBags {
		file, err := os.Open("bags.txt")
		check(err)
		scanner := bufio.NewScanner(file)

		numberOfBags = len(correctBags)
		for scanner.Scan() {
			bags := strings.Fields(scanner.Text())

			mainBag := bags[0] + " " + bags[1]

			if _, exists := correctBags[mainBag]; !exists {
				for i := 5; i < len(bags); i += 4 {
					subBag := bags[i] + " " + bags[i+1]

					if _, exists := correctBags[subBag]; exists || subBag == "shiny gold" {
						correctBags[mainBag] = struct{}{}
						break
					}
				}
			}
		}

		file.Close()
	}

	fmt.Println("Part One:", numberOfBags)
	return
}

func partTwo() {
	file, err := os.Open("bags.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bags bags

	for scanner.Scan() {
		bags = append(bags, scanner.Text())
	}

	fmt.Println("Part Two:", bags.countBags("shiny gold"))
	return
}

func (bags bags) countBags(bag string) int {
	numberOfBags := 0

	for _, bagString := range bags {
		if strings.HasPrefix(bagString, bag) {
			bagFields := strings.Fields(bagString)

			if bagFields[4] != "no" {
				for i := 4; i < len(bagFields); i += 4 {
					numberOfSubBags, err := strconv.Atoi(bagFields[i])
					check(err)

					numberOfBags += numberOfSubBags + numberOfSubBags*bags.countBags(bagFields[i+1]+" "+bagFields[i+2])
				}
			}

			break
		}
	}

	return numberOfBags
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
