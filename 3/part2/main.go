package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	file_content := string(data)

	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllString(file_content, -1)

	enabled := true
	sum := 0

	for _, element := range matches {
		switch {
		case element == "do()":
			enabled = true
		case element == "don't()":
			enabled = false
		case element[:3] == "mul":
			if enabled {
				numRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
				nums := numRe.FindAllStringSubmatch(element, -1)

				for _, match := range nums {
					num1, _ := strconv.Atoi(match[1])
					num2, _ := strconv.Atoi(match[2])
					result := num1 * num2
					sum += result

				}
			}

		}
	}

	fmt.Printf("The result is %d", sum)
}
