package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func main() {
	var leftArray []int
	var rightArray []int
	var sum float64

	leftArray, rightArray = readInput("../2024_1_1.txt")

	sort.Ints(leftArray)
	sort.Ints(rightArray)

	for i := 0; i < len(leftArray); i++ {
		sum += math.Abs(float64(leftArray[i] - rightArray[i]))
	}

	fmt.Printf("%.0f", sum)
}
