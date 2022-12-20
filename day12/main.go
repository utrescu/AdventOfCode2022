package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	x int
	y int
}

func (p Position) GetPossibleMoves(mapa Mapa) []Position {
	directions := []Position{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}

	currentLetter := mapa.Get(p)

	results := []Position{}
	for _, direction := range directions {
		newPos := Position{
			x: p.x + direction.x,
			y: p.y + direction.y,
		}

		if newPos.x >= 0 && newPos.x < mapa.width && newPos.y >= 0 && newPos.y < mapa.height {
			// if pos = actual + 1
			nextPos := mapa.Get(newPos)
			if nextPos >= currentLetter || nextPos == currentLetter-1 {
				results = append(results, newPos)
			}
		}
	}

	return results
}

type Mapa struct {
	width  int
	height int
	cells  [][]rune
}

func (m Mapa) Get(position Position) rune {
	if m.cells[position.y][position.x] == 'S' {
		return 'a' - 1
	}
	return m.cells[position.y][position.x]
}

func readLines(path string) (Mapa, Position, Position, error) {
	file, err := os.Open(path)
	if err != nil {
		return Mapa{}, Position{}, Position{}, err
	}
	defer file.Close()

	start := Position{x: 0, y: 0}
	end := Position{x: 0, y: 0}

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linearray := []rune(scanner.Text())
		posStart, foundStart, posEnd, foundEnd := HasStartOrEnd(linearray)
		if foundStart {
			start.x = posStart
			start.y = len(lines)
			linearray[posStart] = 'a'
		}
		if foundEnd {
			end.x = posEnd
			end.y = len(lines)
			linearray[posEnd] = 'z'
		}
		lines = append(lines, linearray)
	}

	return Mapa{cells: lines, height: len(lines), width: len(lines[0])}, start, end, scanner.Err()
}

func HasStartOrEnd(line []rune) (start int, hasStart bool, end int, hasEnd bool) {

	for i, r := range line {

		if r == 'S' {
			start = i
			hasStart = true
		}

		if r == 'E' {
			end = i
			hasEnd = true
		}
	}
	return
}

func main() {

	mapa, start, end, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	part1, part2 := Day12(mapa, start, end)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

// --- Reverse atempt ---

func Day12(mapa Mapa, start Position, end Position) (int, int) {

	pending := []Position{end}
	distanceToEnd := map[Position]int{end: 0}

	for len(pending) != 0 {
		current := pending[0]
		currentDistance := distanceToEnd[current]
		fmt.Println("Doing: ", string(mapa.Get(current)), " distance:", currentDistance)
		pending = pending[1:]

		for _, candidate := range current.GetPossibleMoves(mapa) {
			_, visited := distanceToEnd[candidate]
			if !visited {
				distanceToEnd[candidate] = currentDistance + 1
				pending = append(pending, candidate)
			}
		}
	}

	minA := distanceToEnd[start]

	for k, v := range distanceToEnd {
		if mapa.Get(k) == 'a' {

			if v < minA {
				minA = v
			}
		}
	}

	return distanceToEnd[start], minA
}
