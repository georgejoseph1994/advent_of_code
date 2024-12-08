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

	i := 0

	for i = 0; i < len(inputArray)-1; i++ {
		if isIncreasing && inputArray[i] > inputArray[i+1] {
			return false
		}

		if !isIncreasing && inputArray[i] < inputArray[i+1] {
			return false
		}

		diff := absInt(inputArray[i] - inputArray[i+1])

		if diff < 1 || diff > 3 {
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

func removeElementFromArray(a []int, i int) []int {
	if len(a) == 0 || i < 0 || i >= len(a) {
		return a
	}

	return append(append([]int{}, a[:i]...), a[i+1:]...)
}

func main() {
	inputArray := readInput("../input.txt")
	var safeArrayCount = 0

	for _, array1 := range inputArray {
		if isSafe(array1) {
			safeArrayCount++
		} else {
			for k := 0; k < len(array1); k++ {
				if isSafe(removeElementFromArray(array1, k)) {
					safeArrayCount++
					break
				}
			}
		}
	}

	fmt.Println(safeArrayCount)

}
