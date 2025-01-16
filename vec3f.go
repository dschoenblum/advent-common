package common

import (
	"fmt"
	"math"
)

type Vec3f struct {
	X, Y, Z float64
}

func NewVec3f(x, y, z float64) Vec3f {
	return Vec3f{x, y, z}
}

func (v Vec3f) Add(other Vec3f) Vec3f {
	return Vec3f{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3f) Sub(other Vec3f) Vec3f {
	return Vec3f{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3f) Scale(multiplier float64) Vec3f {
	return Vec3f{
		X: v.X * multiplier,
		Y: v.Y * multiplier,
		Z: v.Z * multiplier,
	}
}

func (v Vec3f) Dot(other Vec3f) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3f) Cross(other Vec3f) Vec3f {
	return Vec3f{
		X: v.Y*other.Z - v.Z*other.Y,
		Y: v.Z*other.X - v.X*other.Z,
		Z: v.X*other.Y - v.Y*other.X,
	}
}

func (v Vec3f) Divide(divisor float64) Vec3f {
	return Vec3f{
		X: v.X / divisor,
		Y: v.Y / divisor,
		Z: v.Z / divisor,
	}
}

func (v Vec3f) ManhattanDistance() float64 {
	return math.Abs(v.X) + math.Abs(v.Y) + math.Abs(v.Z)
}

func (v Vec3f) String() string {
	return fmt.Sprintf("%f,%f,%f", v.X, v.Y, v.Z)
}
