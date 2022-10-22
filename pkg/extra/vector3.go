package extra

import (
	"fmt"
	"math"
	"reflect"

	"github.com/tch1b0/stargrabber/pkg/general"
)

type Vector3[T general.Number] struct {
	X T
	Y T
	Z T
}

func (v *Vector3[T]) Add(o *Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v *Vector3[T]) Sub(o *Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func (v *Vector3[T]) Mul(o *Vector3[T]) Vector3[T] {
	return Vector3[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
	}
}

func (v *Vector3[T]) ScalarMul(fac T) Vector3[T] {
	return Vector3[T]{
		X: v.X * fac,
		Y: v.Y * fac,
		Z: v.Z * fac,
	}
}

func (v *Vector3[T]) DistanceTo(o *Vector3[T]) float64 {
	return math.Sqrt(
		math.Sqrt(
			math.Pow(float64(o.X-v.X), 2)+
				math.Pow(float64(o.Y-v.Y), 2),
		) + math.Pow(float64(o.Z-v.Z), 2),
	)
}

func (v *Vector3[T]) ManhattanDistanceTo(o *Vector3[T]) float64 {
	return math.Abs(float64(v.X-o.X)) + math.Abs(float64(v.Y-o.Y)) + math.Abs(float64(v.Z-o.Z))
}

func (v *Vector3[T]) String() string {
	var x any = v.X
	var y any = v.Y
	var z any = v.Z

	numType := reflect.TypeOf(v.X).String()

	if numType == "float32" || numType == "float64" {
		return fmt.Sprintf("(%f|%f|%f)", x, y, z)
	} else {
		return fmt.Sprintf("(%d|%d|%d)", x, y, z)
	}
}