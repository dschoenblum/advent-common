package common

import (
	"fmt"
)

type Vec4 struct {
	X, Y, Z, W int
}

var ZeroVec4 = Vec4{}

func NewVec4(x, y, z, w int) Vec4 {
	return Vec4{x, y, z, w}
}

func (v Vec4) Sub(other Vec4) Vec4 {
	return Vec4{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
		W: v.W - other.W,
	}
}

func (v Vec4) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", v.X, v.Y, v.Z, v.W)
}

func (v Vec4) ManhattanDistance() int {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z) + Abs(v.W)
}

func (v Vec4) ManhattanDistanceTo(other Vec4) int {
	return v.Sub(other).ManhattanDistance()
}
