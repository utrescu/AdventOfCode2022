package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	moves, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	tails := []Position{{x: 0, y: 0}}
	fmt.Println("Part 1:", headOcupation(moves, tails))

	for i := 1; i < 9; i++ {
		tails = append(tails, Position{x: 0, y: 0})
	}
	fmt.Println("Part 2:", headOcupation(moves, tails))

}

//  --- Part 1 & 2

func headOcupation(moves []Moves, tails []Position) int {
	positions := map[Position]bool{
		{x: 0, y: 0}: true,
	}

	bridge := Bridge{
		head: Position{x: 0, y: 0},
		tail: tails,
	}

	for _, move := range moves {
		for i := 0; i < move.times; i++ {
			bridge.Move(move.direction)
			last := bridge.tail[len(bridge.tail)-1]
			positions[last] = true
		}
	}

	return len(positions)
}

// --- Position

type Position struct {
	x int
	y int
}

func (p *Position) Move(direction Position) {
	p.x += direction.x
	p.y += direction.y
}

// --- Bridge

type Bridge struct {
	head Position
	tail []Position
}

func (b *Bridge) Move(dir Position) {
	result := []Position{}

	b.head.Move(dir)

	current := b.head
	for i := 0; i < len(b.tail); i++ {
		next := moveTails(current, b.tail[i])
		result = append(result, moveTails(current, next))
		current = next
	}

	b.tail = result
}

func moveTails(head Position, tail Position) Position {
	dx := math.Abs(float64(head.x - tail.x))
	dy := math.Abs(float64(head.y - tail.y))

	newTail := Position{x: tail.x, y: tail.y}

	if dx > 1 || dy > 1 {

		if tail.x != head.x && tail.y != head.y {
			newTail.x += (head.x - tail.x) / int(dx)
			newTail.y += (head.y - tail.y) / int(dy)
		} else {
			if dx > 1 {
				newTail.x += (head.x - tail.x) / int(dx)
			}

			if dy > 1 {
				newTail.y += (head.y - tail.y) / int(dy)
			}
		}
	}
	return newTail
}

// --- Parse file

type Moves struct {
	times     int
	direction Position
}

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func readLines(path string) ([]Moves, error) {

	dirs := map[string]Position{
		"L": {x: -1, y: 0},
		"R": {x: 1, y: 0},
		"U": {x: 0, y: 1},
		"D": {x: 0, y: -1},
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Moves
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		times, _ := stringToInt(parts[1])
		lines = append(lines, Moves{
			direction: dirs[parts[0]],
			times:     times,
		})
	}
	return lines, scanner.Err()
}
