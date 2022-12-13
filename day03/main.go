package main

import (
	"bufio"
	"fmt"
	"os"
)

type rucksack struct {
	all string
	a   string
	b   string
}

func readLines(path string) ([]rucksack, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []rucksack
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mig := len(line) / 2

		first := line[:mig]
		second := line[mig:]

		lines = append(lines, rucksack{all: line, a: first, b: second})
	}
	return lines, scanner.Err()
}

func main() {

	rucksacks, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	fmt.Println("Part 1:", searchRepeated(rucksacks))
	fmt.Println("Part 2:", searchBadges(rucksacks))

}

func removeDuplicateValues(intSlice []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func searchBadge(rucksacks []rucksack) rune {
	keys := map[rune]int{}

	for _, ruck := range rucksacks {
		current := removeDuplicateValues([]rune(ruck.all))
		for _, letter := range current {
			if _, value := keys[letter]; !value {
				keys[letter] = 1
			} else {
				keys[letter] = keys[letter] + 1
				if keys[letter] == 3 {
					return letter
				}
			}
		}
	}
	panic("Inconsistent data! No repetitions")
}

func searchBadges(rucksacks []rucksack) int {
	value := 0
	for i := 0; i < len(rucksacks); i += 3 {
		letter := searchBadge(rucksacks[i : i+3])
		value += getValue(letter)
	}
	return value
}

func searchRepeated(rucksacks []rucksack) int {
	value := 0
	for _, ruck := range rucksacks {
		for _, letter := range ruck.a {
			if contains(letter, ruck.b) {
				value += getValue(letter)
				break
			}
		}
	}
	return value
}

func getValue(letter rune) int {
	if letter > 96 {
		return int(letter-'a') + 1
	} else {
		return int(letter-'A') + 27
	}
}

func contains(letter rune, s string) bool {
	for _, r := range s {
		if r == letter {
			return true
		}
	}
	return false
}
