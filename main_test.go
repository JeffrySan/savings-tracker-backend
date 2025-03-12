package main

import "testing"

var tests = []struct {
	name     string
	dividend float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{"valid data", 100.0, 10.0, 10.0, false},
	{"expect 5", 50, 10.0, 5.0, false},
	{"expect fraction", 51.0, 2.0, 25.5, false},
	{"invalid data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but didnt got one")
			}
			return
		}

		if err != nil {
			t.Error("Expected succes but got an error ", err)
			return
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f ", tt.expected, got)
		}
	}
}
