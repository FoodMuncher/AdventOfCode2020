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

// Please lord, forgive me for this spaghetti code.....

var fields = [7]string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

func partOne() {
	file, err := os.Open("documents.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	document := ""
	validDocuments := 0

	for scanner.Scan() {
		document = document + " " + scanner.Text()

		if scanner.Text() == "" {
			valid := true
			for i := range fields {
				valid = valid && strings.Contains(document, fields[i])
			}

			if valid {
				validDocuments++
			}

			document = ""
		}
	}

	fmt.Println("Part One:", validDocuments)
	return
}

func partTwo() {
	file, err := os.Open("documents.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	document := ""
	validDocuments := 0

	for scanner.Scan() {
		document = document + " " + scanner.Text()

		if scanner.Text() == "" {
			valid := true
			for i := range fields {
				valid = valid && strings.Contains(document, fields[i])
			}

			if valid {
				document := strings.Split(document, " ")

				for _, v := range document {
					if strings.Contains(v, "byr:") {
						id, err := strconv.Atoi(strings.Split(v, ":")[1])
						check(err)

						if !(id >= 1920 && id <= 2002) {
							valid = false
						}

					} else if strings.Contains(v, "iyr:") {
						id, err := strconv.Atoi(strings.Split(v, ":")[1])
						check(err)

						if !(id >= 2010 && id <= 2020) {
							valid = false
						}
					} else if strings.Contains(v, "eyr:") {
						id, err := strconv.Atoi(strings.Split(v, ":")[1])
						check(err)

						if !(id >= 2020 && id <= 2030) {
							valid = false
						}
					} else if strings.Contains(v, "hgt:") {
						height := strings.Split(v, ":")[1]
						if strings.HasSuffix(height, "cm") {
							number, err := strconv.Atoi(strings.Split(height, "cm")[0])
							check(err)
							if !(number >= 150 && number <= 193) {
								valid = false
							}
						} else if strings.HasSuffix(height, "in") {
							number, err := strconv.Atoi(strings.Split(height, "in")[0])
							check(err)
							if !(number >= 59 && number <= 76) {
								valid = false
							}
						} else {
							valid = false
						}
					} else if strings.Contains(v, "hcl:") {
						colour := strings.Split(v, ":")[1]
						if strings.HasPrefix(colour, "#") && len(colour) == 7 {
							for i := 1; i < 7; i++ {
								if !strings.Contains("0123456789abcdef", string(colour[i])) {
									valid = false
								}
							}
						} else {
							valid = false
						}
					} else if strings.Contains(v, "ecl:") {
						colour := strings.Split(v, ":")[1]
						if colour != "amb" && colour != "blu" && colour != "brn" && colour != "gry" && colour != "grn" && colour != "hzl" && colour != "oth" {
							valid = false
						}
					} else if strings.Contains(v, "pid:") {
						if len(strings.Split(v, ":")[1]) != 9 {
							valid = false
						}

					}
				}

				if valid {
					validDocuments++
				}
			}

			document = ""
		}
	}

	fmt.Println("Part One:", validDocuments)
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
