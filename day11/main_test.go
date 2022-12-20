package main

import (
	"day11/monkey"
	"testing"
)

func TestMonkey(t *testing.T) {
	var tests = []struct {
		input          monkey.Monkey
		expectedWorry  int
		expectedMonkey int
	}{
		{
			monkey.Monkey{
				Items:      []int{79, 98},
				Operation:  "*",
				Operators:  []string{"old", "19"},
				Test:       23,
				TestResult: map[bool]int{true: 2, false: 3},
			},
			500,
			3,
		},
		{
			monkey.Monkey{
				Items:      []int{98},
				Operation:  "*",
				Operators:  []string{"old", "19"},
				Test:       23,
				TestResult: map[bool]int{true: 2, false: 3},
			},
			620,
			3,
		},
	}

	for _, tt := range tests {
		t.Run("Test i", func(t *testing.T) {
			worry, nextmonkey := tt.input.SendToMonkey(3, 1)
			if worry != tt.expectedWorry || nextmonkey != tt.expectedMonkey {
				t.Errorf("Test failed '(worry:%d, monkey:%d)' is not (worry:%d, monkey:%d)", worry, nextmonkey, tt.expectedWorry, tt.expectedMonkey)
			}
		})
	}
}
