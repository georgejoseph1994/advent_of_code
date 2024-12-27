package main

import (
	"fmt"
	"os"
	"strings"
)

const XMAS = "XMAS"

func checkXmas(grid [][]rune, i int, j int) int {
	if grid[i][j] != 'X' {
		return 0
	}

	directions := [][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}

	countOfXmas := 0

	for _, dir := range directions {
		for xmasLength := range 4 {
			newX := i + (xmasLength * dir[0])
			newY := j + (xmasLength * dir[1])

			if newX < 0 || newY < 0 || newY > len(grid)-1 || newX > len(grid[0])-1 {
				break
			}

			if grid[newX][newY] != rune(XMAS[xmasLength]) {
				break
			}

			if xmasLength == 3 {
				countOfXmas++
			}
		}
	}

	return countOfXmas
}

func main() {
	data, _ := os.ReadFile("../input.txt")
	content := string(data)

	// load data into a 2d matrix
	var grid [][]rune
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	xmasCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			xmasCount += checkXmas(grid, i, j)
		}
	}

	fmt.Print(xmasCount)

}
