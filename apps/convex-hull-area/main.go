package main

import (
	"fmt" // for console output
	"os" // for command line arguments
    "io/ioutil" // for file io
    "encoding/json" // for json parsing

	"github.com/Shamanskiy/compgeo" // computational geometry package
)

func main() {
	// check if the file name is provided (first programm is the program name itself)
	if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

	// read file content into a byte array
    text, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    // slice of 2D points
    var points []compgeo.Point
    // parse json byte array into the slice
    err = json.Unmarshal([]byte(text), &points)
    if err != nil {
        fmt.Println(err)
        return
    }

    // container for the convex hull (with preallocated memory)
	hull := make([]compgeo.Point, 0, len(points))
	// compute the convex hull of the points
	// NOTE: the order of points in the input slice can be changed
	compgeo.ComputeConvexHull(points, &hull)

	// compute the convex hull area using the shoelace/Gauss's area formula
	area := compgeo.Area(hull)
	fmt.Println(area)
}
