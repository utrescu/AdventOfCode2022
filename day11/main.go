package main

import (
	"bufio"
	"day11/monkey"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	monkeys, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	fmt.Println("Part 1:", Day11(monkeys, 3, 20, 1))

	// Tornar a carregar perquè les dades han canviat!
	// --> No ho trovaba :-(
	monkeys, err = readLines("input")
	if err != nil {
		panic("File read failed")
	}

	var modulo int = 1
	for _, m := range monkeys {
		modulo *= m.Test
	}

	fmt.Println("Part 2:", Day11(monkeys, modulo, 10_000, 2))

}

func Day11(monkeys []monkey.Monkey, lower int, rounds int, part int) int {

	activities := make([]int, len(monkeys))

	for round := 0; round < rounds; round++ {
		for monkeyNumber := 0; monkeyNumber < len(monkeys); monkeyNumber++ {
			moreActivities := len(monkeys[monkeyNumber].Items)
			for len(monkeys[monkeyNumber].Items) > 0 {
				worry, receptor := monkeys[monkeyNumber].SendToMonkey(lower, part)

				monkeys[receptor].AddItem(worry)
			}
			activities[monkeyNumber] += moreActivities
		}

	}

	sort.Sort(sort.Reverse(sort.IntSlice(activities)))

	return activities[1] * activities[0]
}

// ------------ CARREGAR DADES -------------------------

func readLines(path string) ([]monkey.Monkey, error) {

	var re [6]*regexp.Regexp

	re[0] = regexp.MustCompile(`Monkey (\d+)`)
	re[1] = regexp.MustCompile(`(?m)  Starting items: (.*)`)
	re[2] = regexp.MustCompile(`  Operation: new = (\S+) (\S) (\S+)`)
	re[3] = regexp.MustCompile(`  Test: divisible by (\d+)`)
	re[4] = regexp.MustCompile(`    If (\D+): throw to monkey (\d+)`)
	re[5] = regexp.MustCompile(`    If (\D+): throw to monkey (\d+)`)

	var monkeys []monkey.Monkey

	file, err := os.Open(path)
	if err != nil {
		return monkeys, err
	}
	defer file.Close()

	monkeyNumber := 0
	dataLine := 0
	currentMonkey := monkey.Monkey{}
	mapa := make(map[bool]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			monkeyNumber++
			dataLine = 0
			mapa = make(map[bool]int)
			currentMonkey = monkey.Monkey{}
		} else {

			match := re[dataLine].FindStringSubmatch(line)

			switch dataLine {
			case 1:
				stringItems := strings.Split(match[1], ", ")
				var values []int

				for _, v := range stringItems {
					values = append(values, monkey.GetNumber(v))
				}
				currentMonkey.Items = values

			case 2:
				currentMonkey.Operation = match[2]
				currentMonkey.Operators = []string{match[1], match[3]}

			case 3:
				currentMonkey.Test = monkey.GetNumber(match[1])

			case 4:
				mapa[true] = monkey.GetNumber(match[2])

			case 5:
				mapa[false] = monkey.GetNumber(match[2])
				currentMonkey.TestResult = mapa
				monkeys = append(monkeys, currentMonkey)
			}
			dataLine++
		}

	}

	return monkeys, scanner.Err()
}
