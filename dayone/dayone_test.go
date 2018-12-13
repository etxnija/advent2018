package main

import (
	"testing"
)

func TestFrequencyString(t *testing.T) {

	var tests = []struct {
		in     string
		result int
	}{
		{"+1, +1, +1", 3},
		{"+1, +1, -2", 0},
		{"-1, -2, -3", -6},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			r := Frequency(FrequencyString(tt.in))
			if r != tt.result {
				t.Errorf("got %d, but wanted %d", r, tt.result)
			}
		})
	}
}

func TestCallibration(t *testing.T) {

	var tests = []struct {
		in     string
		result int
	}{
		{"+1, -1", 0},
		{"3, +3, +4, -2, -4", 10},
		{"-6, +3, +8, +5, -6", 5},
		{"+7, +7, -2, -7, -4", 14},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			r := Callibration(FrequencyString(tt.in))
			if r != tt.result {
				t.Errorf("got %d, but wanted %d", r, tt.result)
			}
		})
	}
}
