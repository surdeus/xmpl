package vector1

type IVector interface {
	XY() (x, y float64)
}

type Vector struct {
	X, Y float64
}

func (v Vector)XY() (x, y float64) {
	return v.X, v.Y
}

func Add(v, a IVector) IVector {
	x1, y1 := v.XY()
	x2, y2 := a.XY()
	return Vector{x1+x2, y1+y2}
}

