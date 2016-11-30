package rectangle

type Rectangle struct {
	X, Y int
}

func (rec Rectangle) Perimeter() int {
	return rec.X*2 + rec.Y*2
}

func (rec Rectangle) Length() int {
	if k := rec.X; k > rec.Y {
		return rec.X
	}
	return rec.Y
}

func (rec Rectangle) Width() int {
	if rec.Y < rec.X {
		return rec.Y
	} else {
		return rec.X
	}
}
