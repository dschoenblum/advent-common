package common

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

func (g *Grid[T]) toOffset(x, y int) int {
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
