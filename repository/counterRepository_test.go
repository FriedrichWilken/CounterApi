package repository

import (
	"testing"
)

func TestIncreaseBy(t *testing.T) {
	var tests = []struct {
		a           int
		b           int
		expectation int
	}{
		{1, 2, 3},
		{6, 6, 12},
		{-5, 5, 0},
		{12, -12, 0},
	}

	for _, test := range tests {
		repo := New(test.a)
		repo.IncreaseBy(test.b)
		actual := repo.Counter
		if actual != test.expectation {
			t.Error("FAILED: expected was {}, actual was {}", test.expectation, actual)
		}
	}
}

func TestDecreaseBy(t *testing.T) {
	var tests = []struct {
		a           int
		b           int
		expectation int
	}{
		{1, 2, -1},
		{6, 6, 0},
		{-5, 5, -10},
		{12, -12, 24},
	}

	for _, test := range tests {
		repo := New(test.a)
		repo.DecreaseBy(test.b)
		actual := repo.Counter
		if actual != test.expectation {
			t.Error("FAILED: expected was {}, actual was {}", test.expectation, actual)
		}
	}
}
