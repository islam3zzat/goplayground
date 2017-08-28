package utils

import (
	"math"
)

const _DELTA = 0.000001

func step(z, x float64) float64 {
	return z - (z*z-x)/(2*z)
}
func NeotainSqrt(x float64) float64 {
	z := 1.0
	for newZ := step(z, x); math.Abs(newZ-z) > _DELTA; {
		z = newZ
		newZ = step(z, x)
	}
	return z
}
