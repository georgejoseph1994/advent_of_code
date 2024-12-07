package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filePath string) [][]int {
	file, _ := os.Open(filePath)
	defer file.Close()

	var returnArray [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		numbers := make([]int, len(parts))

		for i, part := range parts {
			numbers[i], _ = strconv.Atoi(part)
		}

		returnArray = append(returnArray, numbers)
	}

	return returnArray
}

func isSafe(inputArray []int) bool {
	isIncreasing := true

	if inputArray[0] > inputArray[1] {
		isIncreasing = false
	}

	for i := 0; i < len(inputArray)-1; i++ {
		if isIncreasing && inputArray[i] > inputArray[i+1] {
			return false
		}
		if !isIncreasing && inputArray[i] < inputArray[i+1] {
			return false
		}

		diff := absInt(inputArray[i] - inputArray[i+1])

		if  diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func absInt(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func main() {
	var inputArray [][]int
	inputArray = readInput("../input.txt")
	var safeArrayCount = 0

	for _, array1 := range inputArray {
		if isSafe(array1) {
			safeArrayCount++
		}
	}

	fmt.Println(safeArrayCount)

}
