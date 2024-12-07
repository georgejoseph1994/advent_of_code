package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) ([]int, []int) {
	file, _ := os.Open(filePath)
	defer file.Close()

	var leftArray []int
	var rightArray []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		leftArray = append(leftArray, left)
		rightArray = append(rightArray, right)
	}

	return leftArray, rightArray
}

func getFrequencyMap(numberArray []int) map[int]int {
	frequencyMap := make(map[int]int)

	for _, value := range numberArray {
		frequencyMap[value]++
	}

	return frequencyMap
}

func main() {
	var leftArray []int
	var rightArray []int
	var frequencyMap map[int]int

	leftArray, rightArray = readInput("../2024_1_1.txt")
	frequencyMap = getFrequencyMap(rightArray)

	similarity := 0

	for _, value := range leftArray {
		frequency, exists := frequencyMap[value]
		if exists {
			similarity += (value * frequency)
		}
	}

	fmt.Println(similarity)
}
