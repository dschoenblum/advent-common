package common

type Box struct {
	min, max Vec
}

func NewBox(min, max Vec) Box {
	return Box{min, max}
}

func (b *Box) ExpandToInclude(point Vec) {
	b.min.UpdateMin(point)
	b.max.UpdateMax(point)
}

func (b Box) IsOnBorder(point Vec) bool {
	return point.X == b.min.X || point.X == b.max.X ||
		point.Y == b.min.Y || point.Y == b.max.Y
}

func (b Box) Width() int {
	return b.max.X - b.min.X + 1
}

func (b Box) Height() int {
	return b.max.Y - b.min.Y + 1
}

func (b Box) PointAtOffset(x, y int) Vec {
	return NewVec(b.min.X+x, b.min.Y+y)
}
