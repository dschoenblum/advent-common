package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Vec struct {
	x, y int
}

var ZeroVec = Vec{}
var MinVec = NewVec(math.MinInt, math.MinInt)
var MaxVec = NewVec(math.MaxInt, math.MaxInt)
var North, South, East, West = NewVec(0, 1), NewVec(0, -1), NewVec(1, 0), NewVec(-1, 0)
var Up, Down, Right, Left = North, South, East, West

func NewVec(x, y int) Vec {
	return Vec{x, y}
}

// Create a Vec from "<width>x<height>"
func NewVecFromDimensions(input string) (Vec, error) {
	dimensions := strings.Split(input, "x")
	width, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return ZeroVec, err
	}
	height, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return ZeroVec, err
	}
	return NewVec(width, height), nil
}

func (v Vec) String() string {
	return fmt.Sprintf("{x=%d,y=%d}", v.x, v.y)
}

func (v Vec) X() int {
	return v.x
}

func (v Vec) Y() int {
	return v.y
}

func (v Vec) RotateRight() Vec {
	return Vec{
		x: v.y,
		y: -v.x,
	}
}

func (v Vec) RotateLeft() Vec {
	return Vec{
		x: -v.y,
		y: v.x,
	}
}

func (v Vec) Scale(multiplier int) Vec {
	return Vec{
		x: v.x * multiplier,
		y: v.y * multiplier,
	}
}

func (v Vec) Add(other Vec) Vec {
	return Vec{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vec) Sub(other Vec) Vec {
	return Vec{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

// Return the shortest vector that maintains the same direction
func (v Vec) Shorten() Vec {
	a, b := v.x, v.y
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		a = -a
	}
	return NewVec(v.x/a, v.y/a)
}

func (v Vec) Length() float64 {
	x, y := (float64)(v.x), (float64)(v.y)
	return math.Sqrt(x*x + y*y)
}

func (v Vec) Dot(other Vec) int {
	return v.x*other.x + v.y*other.y
}

func (v Vec) Angle(other Vec) float64 {
	return math.Acos(float64(v.Dot(other)) / (v.Length() * other.Length()))
}

func (v Vec) InSquare(dimensions Vec) bool {
	return 0 <= v.x && v.x < dimensions.x && 0 <= v.y && v.y < dimensions.y
}

func (v *Vec) UpdateMin(other Vec) {
	if other.x < v.x {
		v.x = other.x
	}
	if other.y < v.y {
		v.y = other.y
	}
}

func (v *Vec) UpdateMax(other Vec) {
	if other.x > v.x {
		v.x = other.x
	}
	if other.y > v.y {
		v.y = other.y
	}
}

func (v Vec) ManhattanDistance() int {
	x, y := v.x, v.y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func (v Vec) FindRangeAtY(radius int, y int) (int, int, bool) {
	dist := abs(y - v.y)
	if dist > radius {
		return -1, -1, false
	}
	diff := radius - dist
	minX := v.x - diff
	maxX := v.x + diff
	return minX, maxX, true
}
