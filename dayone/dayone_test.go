package main

import (
	"testing"
)

func TestFrequency(t *testing.T) {

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
			r := Frequency(tt.in)
			if r != tt.result {
				t.Errorf("got %d, but wanted %d", r, tt.result)
			}
		})
	}
}
