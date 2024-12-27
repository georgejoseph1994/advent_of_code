package main

import (
	"fmt"
	"os"
	"strings"
)

const XMAS = "XMAS"

func checkXmas(grid [][]rune, i int, j int) int {
	if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
		return 0
	}

	if grid[i][j] != 'A' {
		return 0
	}

	countOfXmas := 0

	if ((grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S') || (grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M')) &&
		((grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M')) {
		countOfXmas++
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
