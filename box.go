package common

type Box struct {
	min, max Vec
}

func NewBox(min, max Vec) Box {
	return Box{min, max}
}

func NewBoxWithCenter(center Vec, width, height int) Box {
	return Box{
		NewVec(center.X-width/2, center.Y-height/2),
		NewVec(center.X+width/2, center.Y+height/2),
	}
}

func (b *Box) Min() Vec {
	return b.min
}

func (b *Box) Max() Vec {
	return b.max
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

func (b *Box) Translate(amount Vec) {
	b.min = b.min.Add(amount)
	b.max = b.max.Add(amount)
}

func (b *Box) ShiftToHold(pos Vec, margin int) {
	translate := ZeroVec

	leftMargin := pos.X - margin
	if b.min.X > leftMargin {
		translate.X = leftMargin - b.min.X
	} else {
		rightMargin := pos.X + margin
		if b.max.X < rightMargin {
			translate.X = rightMargin - b.max.X
		}
	}

	topMargin := pos.Y - margin
	if b.min.Y > topMargin {
		translate.Y = topMargin - b.min.Y
	} else {
		bottomMargin := pos.Y + margin
		if b.max.Y < bottomMargin {
			translate.Y = bottomMargin - b.max.Y
		}
	}

	if translate != ZeroVec {
		b.Translate(translate)
	}
}

func (b *Box) IsInside(pos Vec) bool {
	return pos.X > b.min.X && pos.X < b.max.X && pos.Y > b.min.Y && pos.Y < b.max.Y
}
