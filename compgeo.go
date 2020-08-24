// Package compgeo implements computational geometry functions
package compgeo

import (
	"math"
	"sort"
	"sync"
)

// eps is a relative tolerance constant for floaint point comparion: abs(a-b) / (abs(a)+abs(b)) < eps
const eps float64 = 1e-17

// Point is a struct respresenting a point in 2D Euclidean space
type Point struct {
	X float64
	Y float64
}

// ======================= Core algorithms ====================//

// ComputeConvexHull computes a convex hull for a given 2D point cloud
func ComputeConvexHull(points []Point, hull *[]Point) {
	// if less then three points provided, return them as a hull
	if len(points) < 3 {
		*hull = append(*hull, points...)
		return
	}

	// sort the points in the hexicographical order X->Y
	sort.Sort(ByXY(points))

	// compute the top half of the convex hull in a separate Goroutine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		computeHalfHull(points, hull, true)
		wg.Done() // signal the main Goroutine when finished
	}()

	// compute the bottom half of the convex hull
	var bottomHull []Point
	computeHalfHull(points, &bottomHull, false)
	wg.Wait() // wait for the top half to finish

	// merge the two halves
	*hull = append(*hull, bottomHull[1:len(bottomHull)-1]...)
}

// computeHalfHull computes onw half of the convex hull.
// If top == true, computes the top half. If top == false, computes the bottom half.
func computeHalfHull(points []Point, hull *[]Point, top bool) {
	N := len(points)
	// add the first two points
	if top {
		*hull = append(*hull, points[0], points[1])
	} else {
		*hull = append(*hull, points[N-1], points[N-2])
	}
	for i := 2; i < len(points); i++ {
		// add the next point
		if top {
			*hull = append(*hull, points[i])
		} else {
			*hull = append(*hull, points[N-i-1])
		}
		// if there are at least 3 points, check the last three for convexity
		for len(*hull) > 2 {
			E := len(*hull) - 1
			// if the last three points form a convex chain, add the next point
			if CheckConvexity((*hull)[E-2], (*hull)[E-1], (*hull)[E]) {
				break
			} else {
				// otherwise, remove the penultimate point, and check the last three again
				*hull = append((*hull)[:E-1], (*hull)[E])
			}
		}
	}
}

// CheckConvexity returns true if points A->B->C form a convex chain.
// The function also takes care of duplicate and nearly identical points
func CheckConvexity(A, B, C Point) bool {
	ABx := B.X - A.X
	ABy := B.Y - A.Y
	BCx := C.X - B.X
	BCy := C.Y - B.Y
	return BCx*ABy-BCy*ABx > 0
}

// Area computes the area of a given polygon using the shoelace formula.
// The polygon is given as a ordered sequence of its vertices
func Area(poly []Point) float64 {
	N := len(poly)
	if N == 0 {
		return 0.
	}
	area := poly[N-1].X*poly[0].Y - poly[0].X*poly[N-1].Y
	for i := 0; i < N-1; i++ {
		area += poly[i].X*poly[i+1].Y - poly[i+1].X*poly[i].Y
	}
	return math.Abs(area) / 2
}

// ======================== Auxiliary =======================//

// ByXY is a wrapper for a slice of Points implementing sort.Interface
type ByXY []Point

// implementation of sotr.Interface for ByXY
func (points ByXY) Len() int      { return len(points) }
func (points ByXY) Swap(i, j int) { points[i], points[j] = points[j], points[i] }
func (points ByXY) Less(i, j int) bool {
	// lexicographics order: first X, then Y
	if CmpFloat(points[i].X, points[j].X) {
		return points[i].Y < points[j].Y
	}
	return points[i].X < points[j].X
}

// CmpFloat performs floating-point comparison with a set relative tolerance.
// It returns true if a==b
// https://floating-point-gui.de/errors/comparison/
func CmpFloat(a float64, b float64) bool {
	absA := math.Abs(a)
	absB := math.Abs(b)
	diff := math.Abs(a - b)

	if absA < math.SmallestNonzeroFloat64 || absB < math.SmallestNonzeroFloat64 || (absA+absB) < math.SmallestNonzeroFloat64 {
		return diff < math.SmallestNonzeroFloat64
	}
	return diff/math.Min(absA+absB, math.MaxFloat64) < eps
}

// ComparePoints performs floating-point comparison for Point objects
func ComparePoints(A, B Point) bool {
	return CmpFloat(A.X, B.X) && CmpFloat(A.Y, B.Y)
}

// ComparePointSets performs floating-point comparison for Point slices
func ComparePointSets(A, B []Point) bool {
	if len(A) != len(B) {
		return false
	}
	for i, v := range A {
		if !ComparePoints(v, B[i]) {
			return false
		}
	}
	return true
}
