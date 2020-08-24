package compgeo

import (
	"math"
	"testing"
)

func TestCheckConvexity(t *testing.T) {

	type test struct {
		pointA Point
		pointB Point
		pointC Point
		answer bool
	}

	tests := []test{
		test{Point{0., 0.}, Point{1., 1.}, Point{2., 1.}, true},
		test{Point{0., 0.}, Point{1., 1.}, Point{0., 1.}, false},
		test{Point{0., 0.}, Point{1., 1.}, Point{0.5, 0.5}, false},
		test{Point{0., 0.}, Point{1., 1.}, Point{1., 1.}, false},
		test{Point{0., 0.}, Point{1., 1.}, Point{0.5, 0.499999999999}, true},
	}

	for i, v := range tests {
		x := CheckConvexity(v.pointA, v.pointB, v.pointC)
		if x != v.answer {
			t.Error("Test", i+1, "Expected", v.answer, "Got", x)
		}
	}
}

func TestCmpFloat(t *testing.T) {

	type test struct {
		x      float64
		y      float64
		answer bool
	}

	tests := []test{
		test{1., 1., true},
		test{2., 1., false},
		test{1., 1.01, false},
		test{1., 1. + eps, true},
		test{1., 1. + 100*eps, false},
		test{math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64, true},
		test{0., 0., true},
	}

	for i, v := range tests {
		x := CmpFloat(v.x, v.y)
		if x != v.answer {
			t.Error("Test", i+1, "Expected", v.answer, "Got", x)
		}
	}
}

func TestArea(t *testing.T) {

	type test struct {
		data   []Point
		answer float64
	}

	tests := []test{
		test{[]Point{Point{0., 0.}, Point{1., 0.}, Point{1., 1.}, Point{0., 1.}}, 1.},
		test{[]Point{Point{0., 0.}, Point{1., 0.}, Point{1., 1.}}, 0.5},
		test{[]Point{Point{0., 0.}, Point{1., 0.}}, 0.},
		test{[]Point{Point{0., 0.}}, 0.},
		test{[]Point{}, 0.},
	}

	for i, v := range tests {
		x := Area(v.data)
		if x != v.answer {
			t.Error("Test", i+1, "Expected", v.answer, "Got", x)
		}
	}
}

func TestComputeConvexHull(t *testing.T) {

	type test struct {
		data   []Point
		answer []Point
	}

	tests := []test{
		test{[]Point{Point{0., 0.}, Point{1., 0.}, Point{1., 1.}, Point{0., 1.}, Point{0.5, 0.5}},
			[]Point{Point{0., 0.}, Point{0., 1.}, Point{1., 1.}, Point{1., 0.}}},
		test{[]Point{Point{0., 0.}, Point{1., 0.}, Point{1., 1.}, Point{0., 0.}, Point{1., 0.}, Point{1., 1.}},
			[]Point{Point{0., 0.}, Point{1., 1.}, Point{1., 0.}}},
	}

	for i, v := range tests {
		x := make([]Point, 0, len(v.data))
		ComputeConvexHull(v.data, &x)
		if !ComparePointSets(x, v.answer) {
			t.Error("Test", i+1, "Expected", v.answer, "Got", x)
		}
	}
}
