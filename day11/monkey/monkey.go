package monkey

import (
	"fmt"
	"strconv"
)

type Monkey struct {
	Items      []int
	Operation  string
	Operators  []string
	Test       int // Test is divisible
	TestResult map[bool]int
}

func (m *Monkey) SendToMonkey(lower int, part int) (worry int, monkey int) {

	// ---- Get Worry level ------------------
	item := m.Items[0]
	m.Items = m.Items[1:]

	op := []int{0, 0}

	for i, operator := range m.Operators {
		if operator == "old" {
			op[i] = item
		} else {
			op[i] = GetNumber(operator)
		}
	}

	switch m.Operation {
	case "*":
		worry = op[0] * op[1]
	case "+":
		worry = op[0] + op[1]
	default:
		panic(fmt.Sprintf("Operation %s not found", m.Operation))
	}
	// ---- Monkey bored
	if part == 1 {
		worry /= lower
	} else {
		worry %= lower
	}

	// ---- Test
	return worry, m.TestResult[worry%m.Test == 0]
}

func (m *Monkey) AddItem(item int) {
	m.Items = append(m.Items, item)
}

func GetNumber(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("%s is not a number", value))
	}
	return number
}
