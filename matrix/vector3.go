package matrix

// Vector3 represents a 3-element vector.
type Vector3 [3]float64

// MulS returns the result of multiplying this vector by a scalar.
func (v Vector3) MulS(s float64) Vector3 {
	return Vector3{v[0] * s, v[1] * s, v[2] * s}
}
