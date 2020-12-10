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
	partOne()
	partTwo()
}

var globalJolts []int
var cache = make(map[int]int)

func partOne() {
	file, err := os.Open("jolts.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var jolts []int

	for scanner.Scan() {
		jolt, err := strconv.Atoi(scanner.Text())
		check(err)
		jolts = append(jolts, jolt)
	}

	sort.Ints(jolts)
	jolts = append([]int{0}, jolts...)           // Adds on the zero jolt beginning socket
	jolts = append(jolts, jolts[len(jolts)-1]+3) // Adds the devices jolts adapter

	oneDiffs := 0
	threesDiffs := 0

	for i := 1; i < len(jolts); i++ {
		diff := jolts[i] - jolts[i-1]

		if diff == 1 {
			oneDiffs++
		} else if diff == 3 {
			threesDiffs++
		}
	}

	fmt.Println("Part One:", oneDiffs*threesDiffs)
	return
}

func partTwo() {
	file, err := os.Open("jolts.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var jolts []int

	for scanner.Scan() {
		jolt, err := strconv.Atoi(scanner.Text())
		check(err)
		jolts = append(jolts, jolt)
	}

	sort.Ints(jolts)
	jolts = append([]int{0}, jolts...)           // Adds on the zero jolt beginning socket
	jolts = append(jolts, jolts[len(jolts)-1]+3) // Adds the devices jolts adapter

	globalJolts = jolts

	fmt.Println("Part Two:", determineCombinations(0))
	return
}

func determineCombinations(startIndex int) int {
	if cachedAnswer, exists := cache[startIndex]; exists {
		return cachedAnswer
	}

	if startIndex == len(globalJolts)-1 {
		cache[startIndex] = 1
		return 1
	}

	combinations := 0
	var endIndex int

	if startIndex+3 > len(globalJolts)-1 {
		endIndex = len(globalJolts) - 1
	} else {
		endIndex = startIndex + 3
	}

	for i := startIndex + 1; i <= endIndex; i++ {
		if globalJolts[i]-globalJolts[startIndex] < 4 {
			combinations += determineCombinations(i)
		} else {
			break
		}
	}

	cache[startIndex] = combinations
	return combinations
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
