package common

import "testing"

func TestGrid_RotateClockwise(t *testing.T) {
	base := NewGridFromString(`abc
def
ghi
jkl`)

	tests := []struct {
		name     string
		rotTimes int
		expect   *Grid[byte]
	}{
		{
			name:     "once",
			rotTimes: 1,
			expect: NewGridFromString(`jgda
kheb
lifc`),
		},
		{
			name:     "twice",
			rotTimes: 2,
			expect: NewGridFromString(`lkj
ihg
fed
cba`),
		},
		{
			name:     "four",
			rotTimes: 4,
			expect:   base, // rotating 4 times returns to original
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := base
			for i := 0; i < tt.rotTimes; i++ {
				have = have.RotateClockwise()
			}

			if !tt.expect.DoesEqual(have) {
				t.Errorf("expect:\n%v\nhave:\n%v", tt.expect, have)
			}
		})
	}
}

func TestGrid_MirrorX(t *testing.T) {
	base := NewGridFromString(`abc
def
ghi
jkl`)
	expect := NewGridFromString(`cba
fed
ihg
lkj`)
	have := base.MirrorX()

	if !expect.DoesEqual(have) {
		t.Errorf("expect:\n%v\nhave:\n%v", expect, have)
	}
}

func TestGrid_MirrorY(t *testing.T) {
	base := NewGridFromString(`abc
def
ghi
jkl`)
	expect := NewGridFromString(`jkl
ghi
def
abc`)
	have := base.MirrorY()

	if !expect.DoesEqual(have) {
		t.Errorf("expect:\n%v\nhave:\n%v", expect, have)
	}
}

func TestGrid_MirrorBoth(t *testing.T) {
	base := NewGridFromString(`abc
def
ghi
jkl`)
	expect := NewGridFromString(`lkj
ihg
fed
cba`)
	have := base.MirrorX().MirrorY()

	if !expect.DoesEqual(have) {
		t.Errorf("expect:\n%v\nhave:\n%v", expect, have)
	}
}
