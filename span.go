package common

type Span struct {
	From, To int
}

func NewSpan(from, to int) Span {
	return Span{from, to}
}

type Spans struct {
	spans []Span
}

func NewSpans() *Spans {
	return &Spans{}
}

func (s *Spans) Count() int {
	c := 0
	for _, s := range s.spans {
		c += s.To - s.From + 1
	}
	return c
}

func (s *Spans) Contains(value int) bool {
	for _, span := range s.spans {
		if value >= span.From && value <= span.To {
			return true
		}
	}
	return false
}

func (s *Spans) AddSpan(span Span) {
	newFrom := span.From
	newTo := span.To

	// Find insertion position: first span whose `to` >= newFrom-1.
	i := 0
	for i < len(s.spans) && s.spans[i].To < newFrom-1 {
		i++
	}

	// Now merge any spans that overlap or touch.
	j := i
	for j < len(s.spans) && s.spans[j].From <= newTo+1 {
		if s.spans[j].From < newFrom {
			newFrom = s.spans[j].From
		}
		if s.spans[j].To > newTo {
			newTo = s.spans[j].To
		}
		j++
	}

	// Replace s[i:j] with the newly merged span.
	merged := Span{From: newFrom, To: newTo}
	s.spans = append(s.spans[:i], append([]Span{merged}, s.spans[j:]...)...)
}
