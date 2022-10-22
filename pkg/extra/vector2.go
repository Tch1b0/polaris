package extra

import (
	"math"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Vector2[T Number] struct {
	X T
	Y T
}

func (v *Vector2[T]) Add(o *Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v *Vector2[T]) Sub(o *Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v *Vector2[T]) Mul(o *Vector2[T]) Vector2[T] {
	return Vector2[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
	}
}

func (v *Vector2[T]) ScalarMul(fac T) Vector2[T] {
	return Vector2[T]{
		X: v.X * fac,
		Y: v.Y * fac,
	}
}

func (v *Vector2[T]) DistanceTo(o *Vector2[T]) float64 {
	return math.Sqrt(math.Pow(float64(o.X-v.X), 2) + math.Pow(float64(o.Y-v.Y), 2))
}

func (v *Vector2[T]) ManhattanDistanceTo(o *Vector2[T]) float64 {
	return math.Abs(float64(v.X-o.X)) + math.Abs(float64(v.Y-o.Y))
}
