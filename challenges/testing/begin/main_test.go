// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"#00", 0},
		{"a2_32_^_&/)", 1},
		{"812_%6ab//", 2},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := letterCounter{identifier: "letter"}.count(tc.input)
			if got != tc.expected {
				t.Errorf("got %v want %v", got, tc.expected)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"#00", 2},
		{"abc_1,?", 1},
		{"abc_12&8_^", 3},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := numberCounter{designation: "number"}.count(tc.input)
			if got != tc.expected {
				t.Errorf("got %v want %v", got, tc.expected)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"#00", 1},
		{"abc_1,?", 3},
		{"abc_12&8_^", 4},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := symbolCounter{}.count(tc.input)
			if got != tc.expected {
				t.Errorf("got %v want %v", got, tc.expected)
			}
		})
	}
}
