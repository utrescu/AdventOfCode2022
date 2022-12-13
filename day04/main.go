package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct {
	start int
	end   int
}

func stringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func NewPair(values string) pair {

	parts := strings.Split(values, "-")

	start, _ := stringToInt(parts[0])
	end, _ := stringToInt(parts[1])
	return pair{start: start, end: end}
}

type assignments struct {
	pairs []pair
}

func (t assignments) CleanContained() bool {
	if t.pairs[0].start <= t.pairs[1].start && t.pairs[0].end >= t.pairs[1].end {
		return true
	}
	if t.pairs[1].start <= t.pairs[0].start && t.pairs[1].end >= t.pairs[0].end {
		return true
	}
	return false
}

func (t assignments) PartialContained() bool {

	line0 := t.pairs[0].start
	line1 := t.pairs[1].start

	for line0 <= t.pairs[0].end && line1 <= t.pairs[1].end {
		if line0 == line1 {
			return true
		} else if line0 < line1 {
			line0++
		} else {
			line1++
		}
	}
	return false
}

func NewAssignments(values string) assignments {
	assignment := assignments{}
	parts := strings.Split(values, ",")
	for _, part := range parts {
		assignment.pairs = append(assignment.pairs, NewPair(part))
	}
	return assignment
}

func readLines(path string) ([]assignments, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []assignments
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		assignment := NewAssignments(line)

		lines = append(lines, assignment)
	}
	return lines, scanner.Err()
}

func main() {

	cleaning, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	part1, part2 := overlapCount(cleaning)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)

}

func overlapCount(cleaningList []assignments) (int, int) {
	count := 0
	count2 := 0
	for _, team := range cleaningList {
		if team.CleanContained() {
			count++
		}

		if team.PartialContained() {
			count2++
		}
	}
	return count, count2
}
