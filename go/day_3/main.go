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

const tree = byte('#')

func partOne() {
	fmt.Println("Part One:", countTrees(3, 1))
	return
}

func partTwo() {
	treeCount := 1

	moves := [5][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for i := 0; i < len(moves); i++ {
		treeCount *= countTrees(moves[i][0], moves[i][1])
	}

	fmt.Println("Part Two:", treeCount)
	return
}

func countTrees(rightMoves, downMoves int) int {
	file, err := os.Open("trees.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	treeCount := 0

	i := 0
	row := 0
	for scanner.Scan() {
		row++
		if row%downMoves == 1 {
			continue
		}

		if scanner.Text()[i] == tree {
			treeCount++
		}

		i = mod(i+rightMoves, 31)
	}

	return treeCount
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
