package matrix

// Matrix3 represents a 3x3 matrix of 3 column vectors.
type Matrix3 [3]Vector3

// Inverse returns the inverse of this matrix, panicking if one does not exist.
func (m Matrix3) Inverse() Matrix3 {
	o := Matrix3{
		{
			m[1][1]*m[2][2] - m[2][1]*m[1][2],
			-(m[0][1]*m[2][2] - m[2][1]*m[0][2]),
			m[0][1]*m[1][2] - m[1][1]*m[0][2],
		},
		{
			-(m[1][0]*m[2][2] - m[2][0]*m[1][2]),
			m[0][0]*m[2][2] - m[2][0]*m[0][2],
			-(m[0][0]*m[1][2] - m[1][0]*m[0][2]),
		},
		{
			m[1][0]*m[2][1] - m[2][0]*m[1][1],
			-(m[0][0]*m[2][1] - m[2][0]*m[0][1]),
			m[0][0]*m[1][1] - m[1][0]*m[0][1],
		},
	}

	det := m[0][0]*o[0][0] + m[1][0]*o[0][1] + m[2][0]*o[0][2]
	if det == 0 {
		panic("matrix is non-invertible")
	}

	o[0][0] /= det
	o[0][1] /= det
	o[0][2] /= det
	o[1][0] /= det
	o[1][1] /= det
	o[1][2] /= det
	o[2][0] /= det
	o[2][1] /= det
	o[2][2] /= det

	return o
}

// MulV returns the result of multiplying this matrix by a vector.
func (m Matrix3) MulV(v Vector3) Vector3 {
	return Vector3{
		m[0][0]*v[0] + m[1][0]*v[1] + m[2][0]*v[2],
		m[0][1]*v[0] + m[1][1]*v[1] + m[2][1]*v[2],
		m[0][2]*v[0] + m[1][2]*v[1] + m[2][2]*v[2],
	}
}
