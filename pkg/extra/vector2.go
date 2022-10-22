package extra

import (
	"fmt"
	"math"
	"reflect"

	"github.com/tch1b0/stargrabber/pkg/general"
)

type Vector2[T general.Number] struct {
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

func (v *Vector2[T]) String() string {
	var x any = v.X
	var y any = v.Y

	numType := reflect.TypeOf(v.X).String()

	if numType == "float32" || numType == "float64" {
		return fmt.Sprintf("(%f|%f)", x, y)
	} else {
		return fmt.Sprintf("(%d|%d)", x, y)
	}
}
