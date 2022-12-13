package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileContent, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	signal := string(fileContent)

	fmt.Println("Part 1:", detectSignal(signal))
	fmt.Println("Part 2:", detectMessage((signal)))

}

func differentContent(messagepart string) bool {
	counts := map[rune]bool{}
	for _, value := range messagepart {
		if _, ok := counts[value]; ok {
			return false
		}
		counts[value] = true
	}
	return true
}

func detectSignal(signal string) int {
	result := 0

	for pos := 4; pos <= len(signal); pos++ {
		if differentContent(signal[pos-4 : pos]) {
			return pos
		}
	}

	return result
}

func detectMessage(signal string) int {
	result := 0

	for pos := 14; pos <= len(signal); pos++ {
		if differentContent(signal[pos-14 : pos]) {
			return pos
		}
	}

	return result
}
