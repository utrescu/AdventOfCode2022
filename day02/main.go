package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type match struct {
	mine  int
	other int
}

func valueMove(play string) int {
	if play == "X" || play == "A" {
		return 1
	} else if play == "Y" || play == "B" {
		return 2
	} else if play == "Z" || play == "C" {
		return 3
	}
	return 0
}

func readLines(path string) ([]match, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []match
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		mine := valueMove(parts[1])
		other := valueMove(parts[0])

		lines = append(lines, match{mine: mine, other: other})
	}
	return lines, scanner.Err()
}

func main() {

	plays, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	fmt.Println("Part 1:", sumPlays(plays))
	fmt.Println("Part 2:", sumPlaysTunned(plays))

}

func calculate(mine int, other int) int {
	if mine == other {
		return 3
	} else if mine == 1 && other == 3 {
		return 6
	} else if mine == 3 && other == 1 {
		return 0
	} else {
		if mine > other {
			return 6
		}
	}
	return 0
}

func sumPlays(matches []match) int {
	sum := 0
	for _, v := range matches {
		sum += v.mine + calculate(v.mine, v.other)
	}
	return sum
}

func sumPlaysTunned(matches []match) int {
	sum := 0
	for _, v := range matches {
		mine := 0
		if v.mine == 1 {
			mine = v.other - 1
			if mine == 0 {
				mine = 3
			}
		} else if v.mine == 3 {
			mine = v.other + 1
			if mine == 4 {
				mine = 1
			}
		} else if v.mine == 2 {
			mine = v.other
		}
		sum += mine + calculate(mine, v.other)
	}
	return sum
}
