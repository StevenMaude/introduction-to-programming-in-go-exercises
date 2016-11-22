package mymath

import "testing"

type testpair struct {
	values   []float64
	expected float64
}

var averageTests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range averageTests {
		v := Average(pair.values)
		if v != pair.expected {
			t.Error(
				"For", pair.values,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

var minTests = []testpair{
	{[]float64{1, 1, 1, 1}, 1},
	{[]float64{5, 7, 6, 4}, 4},
	{[]float64{2, -1, 8, 4}, -1},
	{[]float64{999, 0, 6, 100}, 0},
}

func TestMin(t *testing.T) {
	for _, pair := range minTests {
		m := Min(pair.values)
		if m != pair.expected {
			t.Error("For", pair.values,
				"expected", pair.expected,
				"got", m,
			)
		}
	}
}

var maxTests = []testpair{
	{[]float64{1, 1, 1, 1}, 1},
	{[]float64{5, 7, 6, 4}, 7},
	{[]float64{2, -1, 8, 4}, 8},
	{[]float64{999, 0, 6, 100}, 999},
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		m := Max(pair.values)
		if m != pair.expected {
			t.Error("For", pair.values,
				"expected", pair.expected,
				"got", m,
			)
		}
	}
}
