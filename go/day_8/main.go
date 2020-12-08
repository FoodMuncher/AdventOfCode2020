package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const notDiverged = 0
const diverged = 1
const restarted = 2

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.Open("code.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var code []string

	for scanner.Scan() {
		code = append(code, scanner.Text())
	}

	accumulator := 0
	executedLines := make(map[int]struct{})

	for i := 0; i < len(code); i++ {
		if _, exists := executedLines[i]; exists {
			break
		}

		executedLines[i] = struct{}{}

		action := code[i][:3]
		number, err := strconv.Atoi(code[i][4:])
		check(err)

		switch action {
		case "acc":
			accumulator += number
		case "jmp":
			i += number - 1
		case "nop":
			continue
		}
	}

	fmt.Println("Part One:", accumulator)
	return
}

func partTwo() {
	file, err := os.Open("code.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var code []string

	for scanner.Scan() {
		code = append(code, scanner.Text())
	}

	accumulator := 0
	executedLines := make(map[int]struct{})

	var divergedIndex int
	divergedStatus := notDiverged
	divergedExecutedLines := make(map[int]struct{})
	divergedAccumulator := 0

	for i := 0; i < len(code); i++ {
		if _, exists := executedLines[i]; exists {
			i = divergedIndex - 1
			executedLines = copyMap(divergedExecutedLines)
			divergedStatus = restarted
			accumulator = divergedAccumulator
			continue
		}

		action := code[i][:3]
		number, err := strconv.Atoi(code[i][4:])
		check(err)

		executedLines[i] = struct{}{}

		switch action {
		case "acc":
			accumulator += number
		case "jmp":
			if divergedStatus == notDiverged {
				divergedStatus = diverged
				divergedAccumulator = accumulator
				divergedIndex = i
				delete(executedLines, i)
				divergedExecutedLines = copyMap(executedLines)
				executedLines[i] = struct{}{}
			} else {
				i += number - 1
			}
		case "nop":
			if divergedStatus == notDiverged {
				divergedStatus = diverged
				divergedAccumulator = accumulator
				divergedIndex = i
				delete(executedLines, i)
				divergedExecutedLines = copyMap(executedLines)
				executedLines[i] = struct{}{}
				i += number - 1
			}
		}

		if divergedStatus == restarted {
			divergedStatus = notDiverged
		}
	}

	fmt.Println("Part Two:", accumulator)
	return
}

func copyMap(original map[int]struct{}) map[int]struct{} {
	new := make(map[int]struct{})

	for key := range original {
		new[key] = struct{}{}
	}

	return new
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
