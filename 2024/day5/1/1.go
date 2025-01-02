package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([][]int, [][]int) {
	data, _ := os.ReadFile("../input.txt")
	content := string(data)

	var rules [][]int
	var updates [][]int

	lines := strings.Split(content, "\n")
	emptyLineEncountered := false

	for _, line := range lines {
		if line == "" {
			emptyLineEncountered = true
			continue
		}

		if !emptyLineEncountered {
			parts := strings.Split(line, "|")
			val0, _ := strconv.Atoi(parts[0])
			val1, _ := strconv.Atoi(parts[1])
			row := []int{val0, val1}
			rules = append(rules, row)
		} else {
			parts := strings.Split(line, ",")

			var row []int

			for _, p := range parts {
				val, _ := strconv.Atoi(p)
				row = append(row, val)
			}
			updates = append(updates, row)
		}
	}

	return rules, updates
}

func checkRule(rules [][]int, updates [][]int) {
	var rejectRow bool
	sum := 0

	for _, updatesRow := range updates {
		rejectRow = false

		for _, rulesRow := range rules {
			isFirstFound := false
			isSecondFound := false

			for _, u := range updatesRow {
				isFirstFound = (rulesRow[0] == u)
				if isFirstFound && isSecondFound {
					rejectRow = true
					break
				}
				isSecondFound = (rulesRow[1] == u)
			}

			if rejectRow {
				break
			}
		}

		if !rejectRow {
			sum += updatesRow[len(updatesRow)/2]
		}
	}
	fmt.Println(sum)
}

func main() {
	rules, updates := parseInput()

	checkRule(rules, updates)
}
