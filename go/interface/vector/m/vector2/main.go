package vector2

import(
	"vector/m/vector1"
)

type Vector vector1.Vector
type IVector vector1.IVector

func (v Vector)XY() (x, y float64) {
	return v.X, v.Y
}

func Mul(v IVector, m float64) IVector {
	x, y := v.XY()
	return Vector{x*m, y*m}
}

