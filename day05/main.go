package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	bigcargo, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	fmt.Println("Part 1:", bigcargo.Move())

	bigcargo, err = readLines("input")
	if err != nil {
		panic("File read failed")
	}
	fmt.Println("Part 2:", bigcargo.Move9001())

}

type move struct {
	quantity int
	from     int
	to       int
}

type cargo struct {
	board map[int][]string

	movements []move
}

func (c cargo) generateResult() string {
	result := ""

	keys := make([]int, 0, len(c.board))
	for k := range c.board {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		result += c.board[key][0]
	}
	return result
}

func (c cargo) Move() string {

	for _, change := range c.movements {
		for i := 0; i < change.quantity; i++ {
			fromvalues := c.board[change.from]
			tovalues := append([]string{fromvalues[0]}, c.board[change.to]...)
			c.board[change.to] = tovalues
			c.board[change.from] = fromvalues[1:]
		}
	}
	return c.generateResult()
}

func (c cargo) Move9001() string {

	for _, change := range c.movements {
		containersmoved := make([]string, change.quantity)
		fromvalues := c.board[change.from]
		copy(containersmoved, fromvalues[0:change.quantity])
		tovalues := append(containersmoved, c.board[change.to]...)
		c.board[change.to] = tovalues
		c.board[change.from] = fromvalues[change.quantity:]
	}
	return c.generateResult()
}

/// ----- Only for LOAD data  -------------------------------

func StringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func lineContains(line string, char string) bool {
	for _, b := range line {
		if string(b) == "[" {
			return true
		}
	}
	return false
}

func readLines(path string) (cargo, error) {
	var re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	result := cargo{}

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	loadingBoard := true
	board := map[int][]string{}
	moves := []move{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			loadingBoard = false
			continue
		}

		if loadingBoard {
			pos := 1
			if !lineContains(line, "[") {
				continue
			}
			for i := 1; i < len(line); i += 4 {
				caracter := string(line[i])
				if caracter != " " {
					if value, exists := board[pos]; !exists {
						board[pos] = []string{caracter}
					} else {
						board[pos] = append(value, caracter)
					}
				}
				pos++

			}
		} else {
			match := re.FindStringSubmatch(line)
			quantity, _ := StringToInt(match[1])
			from, _ := StringToInt(match[2])
			to, _ := StringToInt(match[3])
			moves = append(moves, move{quantity: quantity, from: from, to: to})
		}

	}
	result.board = board
	result.movements = moves
	return result, scanner.Err()
}
