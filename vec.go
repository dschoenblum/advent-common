package common

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Vec struct {
	X, Y int
}

var ZeroVec = Vec{}
var MinVec = NewVec(math.MinInt, math.MinInt)
var MaxVec = NewVec(math.MaxInt, math.MaxInt)
var North, South, East, West = NewVec(0, 1), NewVec(0, -1), NewVec(1, 0), NewVec(-1, 0)
var Up, Down, Right, Left = North, South, East, West
var Directions = []Vec{Up, Down, Left, Right}

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
	return fmt.Sprintf("{x=%d,y=%d}", v.X, v.Y)
}

func (v Vec) Csv() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

func (v Vec) RotateRight() Vec {
	return Vec{
		X: v.Y,
		Y: -v.X,
	}
}

func (v Vec) RotateLeft() Vec {
	return Vec{
		X: -v.Y,
		Y: v.X,
	}
}

func (v Vec) Reverse() Vec {
	return Vec{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vec) Scale(multiplier int) Vec {
	return Vec{
		X: v.X * multiplier,
		Y: v.Y * multiplier,
	}
}

func (v Vec) Add(other Vec) Vec {
	return Vec{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec) Sub(other Vec) Vec {
	return Vec{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Return the shortest vector that maintains the same direction
func (v Vec) Shorten() Vec {
	a, b := v.X, v.Y
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		a = -a
	}
	return NewVec(v.X/a, v.Y/a)
}

func (v Vec) Length() float64 {
	x, y := (float64)(v.X), (float64)(v.Y)
	return math.Sqrt(x*x + y*y)
}

func (v Vec) Dot(other Vec) int {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec) Angle(other Vec) float64 {
	return math.Acos(float64(v.Dot(other)) / (v.Length() * other.Length()))
}

func (v Vec) InSquare(dimensions Vec) bool {
	return 0 <= v.X && v.X < dimensions.X && 0 <= v.Y && v.Y < dimensions.Y
}

func (v *Vec) UpdateMin(other Vec) {
	if other.X < v.X {
		v.X = other.X
	}
	if other.Y < v.Y {
		v.Y = other.Y
	}
}

func (v *Vec) UpdateMax(other Vec) {
	if other.X > v.X {
		v.X = other.X
	}
	if other.Y > v.Y {
		v.Y = other.Y
	}
}

func (v Vec) ManhattanDistance() int {
	x, y := v.X, v.Y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func (v Vec) FindRangeAtY(radius int, y int) (int, int, bool) {
	dist := Abs(y - v.Y)
	if dist > radius {
		return -1, -1, false
	}
	diff := radius - dist
	minX := v.X - diff
	maxX := v.X + diff
	return minX, maxX, true
}
