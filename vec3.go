package common

import (
	"fmt"
	"math"
)

type Vec3 struct {
	X, Y, Z int
}

var ZeroVec3 = Vec3{}
var MinVec3 = NewVec3(math.MinInt, math.MinInt, math.MinInt)
var MaxVec3 = NewVec3(math.MaxInt, math.MaxInt, math.MaxInt)

func NewVec3(x, y, z int) Vec3 {
	return Vec3{x, y, z}
}

func (v Vec3) String() string {
	return fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z)
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v *Vec3) AddSelf(other Vec3) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3) ManhattanDistance() int {
	return Abs(v.X) + Abs(v.Y) + Abs(v.Z)
}

func (v Vec3) ManhattanDistanceTo(other Vec3) int {
	return v.Sub(other).ManhattanDistance()
}

func (v Vec3) DistanceSquared() int {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec3) UpdateMin(other Vec3) {
	if other.X < v.X {
		v.X = other.X
	}
	if other.Y < v.Y {
		v.Y = other.Y
	}
	if other.Z < v.Z {
		v.Z = other.Z
	}
}

func (v *Vec3) UpdateMax(other Vec3) {
	if other.X > v.X {
		v.X = other.X
	}
	if other.Y > v.Y {
		v.Y = other.Y
	}
	if other.Z > v.Z {
		v.Z = other.Z
	}
}
