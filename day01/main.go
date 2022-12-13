package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func stringToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

func readLines(path string) ([]int, error) {
	max := []int{0, 0, 0}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	current := 0
	for scanner.Scan() {

		line := scanner.Text()
		if len(line) == 0 {
			max = insertIfMax(max, current)
			current = 0
		} else {
			number, _ := stringToInt(line)
			current += number
		}
	}

	max = insertIfMax(max, current)

	return max, scanner.Err()
}

func insertIfMax(maxarray []int, current int) []int {
	newArray := append(maxarray, current)
	sort.Sort(sort.Reverse(sort.IntSlice(newArray)))
	return newArray[0:3]
}

func main() {

	argLength := len(os.Args[1:])
	if argLength != 1 {
		fmt.Println("Has de passar el nom del fitxer com a paràmetre")
		os.Exit(1)
	}

	max, err := readLines(os.Args[1])
	if err != nil {
		panic("File read failed")
	}

	fmt.Println("Part 1:", max[0])

	sum := 0
	for _, v := range max {
		sum += v
	}
	fmt.Println("Part 2:", sum)

}
