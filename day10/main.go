package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	instructions, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	part1 := Part1(instructions, 20, 40)

	fmt.Println("Part 1:", part1)

	part2 := Part2(instructions, 40)

	for _,v := range part2 {
		fmt.Println(v)
	}
}

func Part1(instructions []instruction, firstCicle int, increments int) int {
	result := 0
	x := 1
	cicle := 0
	cicleResultAt := firstCicle
	for _, current := range instructions {
		// fmt.Println(current.operation, current.argument)
		for i := 0; i < current.cicles; i++ {
			cicle++

			// fmt.Println("cicle", cicle, "x:", x)
			if cicle == cicleResultAt {
				// fmt.Println(cicle, "(", x, ")=", cicle*x)
				result += cicle * x
				cicleResultAt += increments
			}

			if i == 1 {
				x += current.argument
			}
			if cicle == 220 {
				return result
			}
		}
	}
	return result
}

func Part2(instructions []instruction, increments int) []string {
	result := []string{}
	x := 1
	cicle := 0

	line := ""
	for _, current := range instructions {
		// fmt.Println(current.operation, current.argument)
		for i := 0; i < current.cicles; i++ {
			cicle++
			if len(line) == x || len(line) == x-1 || len(line) == x+1 {
				line += "#"
			} else {
				line += "."
			}

			// fmt.Println("cicle", cicle, "x:", x)
			if cicle == increments {

				result = append(result, line)
				// fmt.Println(cicle, "(", x, ")=", cicle*x)
				if len(result) == 6 {
					return result
				}
				line = ""
				cicle = 0
			}

			if i == 1 {
				x += current.argument
			}

		}
	}
	return result
}

func readLines(path string) ([]instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		argument := ""
		if len(parts) > 1 {
			argument = parts[1]
		}

		value := NewInstruction(parts[0], argument)

		lines = append(lines, value)
	}
	return lines, scanner.Err()
}

func stringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

type instruction struct {
	operation string
	cicles    int
	argument  int
}

func NewInstruction(operation string, argument string) instruction {
	if operation == "noop" {
		return instruction{operation: operation, cicles: 1, argument: 0}
	} else {
		value, _ := stringToInt(argument)
		return instruction{operation: operation, cicles: 2, argument: value}
	}
}
