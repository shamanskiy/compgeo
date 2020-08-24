# compgeo
This is a small programming exercise in [Go](https://golang.org/) language. 
Its goal is to implement several computational geometry algorithms from [this book](https://www.springer.com/gp/book/9783540779735) by De Berg et al.

## How to build
You can use Go built-in tools to download, build and install this project, all in one go! 

#### Package
To get only the library, a.k.a. Go package, do
```
go get github.com/Shamanskiy/compgeo
```
This will clone the `github.com/Shamanskiy/compgeo.git` repository to `GO_PATH/src/github.com/Shamanskiy/compgeo`, build the `compgeo` package and install it to `GO_PATH/pkg/YOUR_SYSTEM/github.com/Shamanskiy`. Now your application can use `compgeo` by importing it in this way:
```
import (
    "github.com/Shamanskiy/compgeo"
)
```

#### Examples
If you want to test the examples included in `compgeo`, do
```
go get github.com/Shamanskiy/compgeo/apps/APP-NAME
```
This will build `APP-NAME` and install it to `GO_PATH/bin`.


## List of available algorithms
So far, the list of available algorithms is rather humble. It includes:
* [Graham's scan](https://en.wikipedia.org/wiki/Graham_scan) for computing the convex hull of a point cloud
* [the shoelace algorithm](https://en.wikipedia.org/wiki/Shoelace_formula) for computing the area of a simple polygon

## List of examples
All examples require a JSON file with points as an input. For example, the following file defines a list of 3 points in 2D space:
```
[
{"x": 0,"y": 0},
{"x": 1,"y": 0},
{"x": 0,"y": 1}
]
```
The points are (0,0), (1,0) and (0,1). Points in 3D space are encoded in a similar fashion with an extra `z` field. The `data` folder in the `compgeo.git` repository contains several input files to use with the examples.

#### convex-hull-area
In this example, we compute the convex hull of a point cloud and compute its area. You can run it for the `diamond.json` point cloud by executing
```
GO_PATH/bin/convex-hull-area COMPGEO_PATH/data/diamond.json
```
The result `2` will be printed to the console.

## Testing
There are some unit tests included in the `compgeo` project. You can run the tests by typing
```
go test [-v] github.com/Shamanskiy/compgeo
```
Optional flag `-v` increases the verbosity level of the tests.
