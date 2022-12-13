package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func stringArrayToInt(stringArray []string) ([]int, error) {
	var result []int
	for _, value := range stringArray {
		numero, err := stringToInt(value)
		if err != nil {
			return nil, err
		}
		result = append(result, numero)
	}
	return result, nil
}

func readLines(path string) (Zone, error) {
	file, err := os.Open(path)
	if err != nil {
		return Zone{}, err
	}
	defer file.Close()

	var zone [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		caracters := strings.Split(scanner.Text(), "")
		trees, _ := stringArrayToInt(caracters)
		zone = append(zone, trees)
	}
	return Zone{cells: zone}, scanner.Err()
}

func main() {
	zone, err := readLines("input")
	if err != nil {
		panic("File read failed")
	}

	count, space := countVisibleTrees(zone)
	fmt.Println("Part 1:", count)
	fmt.Println("Part 2:", space)

}

type Position struct {
	x int
	y int
}

func (p *Position) Move(dir Position) {
	p.x += dir.x
	p.y += dir.y
}

type Zone struct {
	cells [][]int
}

func (z Zone) Height() int {
	return len(z.cells)
}

func (z Zone) Width() int {
	return len(z.cells[0])
}

func (m Zone) Get(p Position) int {
	return m.cells[p.y][p.x]
}

func (m Zone) IsInside(p Position) bool {

	if p.x >= 0 && p.x < len(m.cells[0]) && p.y >= 0 && p.y < len(m.cells) {
		return true
	}

	return false
}

// ---- PART 1

func countVisibleTrees(zone Zone) (int, int) {
	count := 0
	space := 0

	for posy := 0; posy < zone.Width(); posy++ {
		for posx := 0; posx < zone.Height(); posx++ {
			visible, newspace := isVisible(zone, Position{x: posx, y: posy})
			count += visible
			if newspace > space {
				space = newspace
			}
		}
	}
	return count, space
}

func isVisible(zone Zone, position Position) (int, int) {
	directions := []Position{
		{x: 0, y: 1},
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: -1, y: 0},
	}

	treeTall := zone.Get(position)

	// Mirar si és visible i si ho és retorna 1
	failed := 0
	vision := 1
	for _, direction := range directions {
		current := Position{x: position.x, y: position.y}
		visiondirection := 0
		for zone.IsInside(current) {
			current.Move(direction)
			if zone.IsInside(current) {
				visiondirection++
				if treeTall <= zone.Get(current) {
					failed++
					break
				}

			} else {
				break
			}
		}
		vision *= visiondirection
	}

	if failed < 4 {
		// fmt.Println("Found", position)
		return 1, vision
	}
	return 0, vision
}
