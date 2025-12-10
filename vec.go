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
var (
	South     = NewVec(0, 1)
	East      = NewVec(1, 0)
	West      = NewVec(-1, 0)
	North     = NewVec(0, -1)
	NorthEast = NewVec(1, -1)
	NorthWest = NewVec(-1, -1)
	SouthEast = NewVec(1, 1)
	SouthWest = NewVec(-1, 1)
)
var (
	Down      = South
	Right     = East
	Left      = West
	Up        = North
	UpRight   = NorthEast
	UpLeft    = NorthWest
	DownRight = SouthEast
	DownLeft  = SouthWest
)
var Directions = []Vec{Up, Down, Left, Right}

func NewVec(x, y int) Vec {
	return Vec{x, y}
}

// NewVecFromDimensions creates a Vec from "<width>x<height>"
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

// NewVecFromCSV creates a Vec from "<x>,<y>"
func NewVecFromCSV(input string) (Vec, error) {
	dimensions := strings.Split(input, ",")
	x, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return ZeroVec, err
	}
	y, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return ZeroVec, err
	}
	return NewVec(x, y), nil
}

func (v Vec) String() string {
	//return fmt.Sprintf("{x=%d,y=%d}", v.X, v.Y)
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

func (v Vec) Csv() string {
	return fmt.Sprintf("%d,%d", v.X, v.Y)
}

func (v Vec) RotateRight() Vec {
	return Vec{
		X: -v.Y,
		Y: v.X,
	}
}

func (v Vec) RotateLeft() Vec {
	return Vec{
		X: v.Y,
		Y: -v.X,
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

func (v Vec) UpdateMinMax(min, max *Vec) {
	min.UpdateMin(v)
	max.UpdateMax(v)
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

func (v Vec) Clamp(min, max Vec) Vec {
	if v.X < min.X {
		v.X = min.X
	}
	if v.X > max.X {
		v.X = max.X
	}
	if v.Y < min.Y {
		v.Y = min.Y
	}
	if v.Y > max.Y {
		v.Y = max.Y
	}
	return v
}
