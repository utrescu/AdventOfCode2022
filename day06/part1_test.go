package main

import "testing"

func TestDifferent(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abab", false},
		{"abca", false},
		{"abcb", false},
		{"abcc", false},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			result := differentContent(tt.input)
			if result != tt.expected {
				t.Errorf("Test failed '%s'", tt.input)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"dfdaabcd", 8},
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			result := detectSignal(tt.input)
			if result != tt.expected {
				t.Errorf("Test failed '%s', wants %d and received %d", tt.input, tt.expected, result)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tt := range tests {
		testname := tt.input
		t.Run(testname, func(t *testing.T) {
			result := detectMessage(tt.input)
			if result != tt.expected {
				t.Errorf("Test failed '%s', wants %d and received %d", tt.input, tt.expected, result)
			}
		})
	}
}
