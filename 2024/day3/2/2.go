package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const DONT = "don't()"
const DO = "do()"

func extractMulNumbers(inputSting string) [][]string {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	matches := re.FindAllStringSubmatch(inputSting, -1)

	var result [][]string
	for _, match := range matches {
		if match[0] == DONT || match[0] == DO {
			result = append(result, []string{match[0]})
			continue
		}

		if len(match) == 3 {
			result = append(result, []string{match[1], match[2]})
		}
	}
	return result
}

func main() {
	data, _ := os.ReadFile("../input.txt")
	input := string(data)
	dataArray := extractMulNumbers(input)
	result := 0
	isCounting := true

	for _, row := range dataArray {

		if row[0] == DONT {
			isCounting = false
			continue
		} else if row[0] == DO {
			isCounting = true
			continue
		}

		if isCounting {
			x, _ := strconv.Atoi(row[0])
			y, _ := strconv.Atoi(row[1])
			result += x * y
		}
	}

	fmt.Println(result)
}
