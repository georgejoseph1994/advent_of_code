package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func extractMulNumbers(inputSting string) [][]int {

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(inputSting, -1)

	var result [][]int
	for _, match := range matches {
		if len(match) == 3 {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			result = append(result, []int{x, y})
		}
	}

	return result
}

func main() {
	data, _ := os.ReadFile("../input.txt")
	input := string(data)
	dataArray := extractMulNumbers(input)
	result:=0

	for _,row := range(dataArray){
		result+= row[0] * row[1]
	}

	fmt.Println(result)

}
