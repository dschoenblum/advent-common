package common

import (
	"iter"
)

type Grid[T comparable] struct {
	width, height int
	values        []T
}

func NewGrid[T comparable](width, height int) *Grid[T] {
	return &Grid[T]{
		width:  width,
		height: height,
		values: make([]T, width*height),
	}
}

func NewGridFromInput[T comparable](input string, parse func(rune) T) *Grid[T] {
	lines := ToTrimmedLines(input)
	grid := NewGrid[T](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, r := range line {
			grid.Set(x, y, parse(r))
		}
	}
	return grid
}

func NewGridFromString(input string) *Grid[byte] {
	return NewGridFromInput(input, func(r rune) byte {
		return byte(r)
	})
}

func (g *Grid[T]) toOffset(x, y int) int {
	return y*g.width + x
}

func (g *Grid[T]) toOffsetRepeating(x, y int) int {
	x = x % g.width
	if x < 0 {
		x += g.width
	}
	y = y % g.height
	if y < 0 {
		y += g.height
	}
	return y*g.width + x
}

func (g *Grid[T]) Width() int {
	return g.width
}

func (g *Grid[T]) Height() int {
	return g.height
}

func (g *Grid[T]) SetVec(pos Vec, value T) {
	g.Set(pos.X, pos.Y, value)
}

func (g *Grid[T]) Set(x, y int, value T) {
	g.values[g.toOffset(x, y)] = value
}

func (g *Grid[T]) GetVec(pos Vec) (value T) {
	return g.Get(pos.X, pos.Y)
}

func (g *Grid[T]) Get(x, y int) (value T) {
	return g.values[g.toOffset(x, y)]
}

func (g *Grid[T]) GetVecRepeating(pos Vec) (value T) {
	return g.GetRepeating(pos.X, pos.Y)
}

func (g *Grid[T]) GetRepeating(x, y int) (value T) {
	return g.values[g.toOffsetRepeating(x, y)]
}

func (g *Grid[T]) TryGetVec(pos Vec) (value T, ok bool) {
	return g.TryGet(pos.X, pos.Y)
}

func (g *Grid[T]) TryGet(x, y int) (value T, ok bool) {
	if !g.IsValid(x, y) {
		return
	}
	return g.values[g.toOffset(x, y)], true
}

func (g *Grid[T]) IsValidVec(pos Vec) bool {
	return g.IsValid(pos.X, pos.Y)
}

func (g *Grid[T]) IsValid(x, y int) bool {
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}

func (g *Grid[T]) DoesEqual(other *Grid[T]) bool {
	if g.width != other.width || g.height != other.height {
		return false
	}
	for i := 0; i < len(g.values); i++ {
		if g.values[i] != other.values[i] {
			return false
		}
	}
	return true
}

// First finds the first position in the grid that satisfies the predicate.
func (g *Grid[T]) First(predicate func(T) bool) (pos Vec, ok bool) {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			if predicate(g.Get(x, y)) {
				return Vec{x, y}, true
			}
		}
	}
	return Vec{}, false
}

// Find finds the first position in the grid that matches the value.
func (g *Grid[T]) Find(value T) (pos Vec, ok bool) {
	return g.First(func(v T) bool {
		return v == value
	})
}

// Visit calls the visitor function for each position in the grid.
func (g *Grid[T]) Visit(visitor func(Vec, T)) {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			visitor(Vec{x, y}, g.Get(x, y))
		}
	}
}

func (g *Grid[T]) All() iter.Seq2[Vec, T] {
	return func(yield func(pos Vec, value T) bool) {
		for y := 0; y < g.height; y++ {
			for x := 0; x < g.width; x++ {
				if !yield(NewVec(x, y), g.Get(x, y)) {
					return
				}
			}
		}
	}
}

var neighbors4 = []Vec{North, East, South, West}

func (g *Grid[T]) Neighbors4(pos Vec) iter.Seq2[Vec, T] {
	return func(yield func(neighbor Vec, value T) bool) {
		for _, dir := range neighbors4 {
			neighbor := pos.Add(dir)
			if g.IsValidVec(neighbor) {
				if !yield(neighbor, g.GetVec(neighbor)) {
					return
				}
			}
		}
	}
}

func (g *Grid[T]) Neighbors4Repeating(pos Vec) iter.Seq2[Vec, T] {
	return func(yield func(neighbor Vec, value T) bool) {
		for _, dir := range neighbors4 {
			neighbor := pos.Add(dir)
			if !yield(neighbor, g.GetVecRepeating(neighbor)) {
				return
			}
		}
	}
}

var neighbors8 = []Vec{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}

func (g *Grid[T]) Neighbors8(pos Vec) iter.Seq2[Vec, T] {
	return func(yield func(neighbor Vec, value T) bool) {
		for _, dir := range neighbors8 {
			neighbor := pos.Add(dir)
			if g.IsValidVec(neighbor) {
				if !yield(neighbor, g.GetVec(neighbor)) {
					return
				}
			}
		}
	}
}

func (g *Grid[T]) RotateClockwise() *Grid[T] {
	n := NewGrid[T](g.height, g.width)
	g.Visit(func(pos Vec, v T) {
		n.Set(g.height-pos.Y-1, pos.X, v)
	})
	return n
}

func (g *Grid[T]) MirrorX() *Grid[T] {
	n := NewGrid[T](g.width, g.height)
	g.Visit(func(pos Vec, v T) {
		n.Set(g.width-pos.X-1, pos.Y, v)
	})
	return n
}

func (g *Grid[T]) MirrorY() *Grid[T] {
	n := NewGrid[T](g.width, g.height)
	g.Visit(func(pos Vec, v T) {
		n.Set(pos.X, g.height-pos.Y-1, v)
	})
	return n
}
