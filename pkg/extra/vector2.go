package extra

import "math"

type Vector2 struct {
	X int
	Y int
}

func (v *Vector2) DistanceTo(other *Vector2) float64 {
	return math.Sqrt(math.Pow(float64(other.X-v.X), 2) + math.Pow(float64(other.Y-v.Y), 2))
}

func (v *Vector2) ManhattanDistanceTo(other *Vector2) float64 {
	return math.Abs(float64(v.X-other.X)) + math.Abs(float64(v.Y-other.Y))
}
