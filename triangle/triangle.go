// Package triangle contains tools for identifying triangle types
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
	Deg             // degenerate / line
)

// KindFromSides returns the type of triangle created by the three inputs, or NaT if the inputs do not form a triangle
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !isATriangle(a, b, c):
		return NaT
	case a+b == c:
		return Deg
	}

	equalSides := bToU(a == b) + bToU(a == c) + bToU(b == c)

	switch equalSides {
	case 0:
		return Sca
	case 1:
		return Iso
	case 3:
		return Equ
	}
	// This should never happen, so if it does we have some weird non-triangle
	return NaT
}

// bToU coerces a boolean to an integer with true being 1 and false being 0
func bToU(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

// isATriangle checks that all sides of a triangle are positive real numbers that fulfill the triangle inequality
func isATriangle(a, b, c float64) bool {
	sides := [3]float64{a, b, c}

	for _, side := range sides {
		switch {
		case side <= 0:
			return false
		case math.IsNaN(side):
			return false
		case math.IsInf(side, 0):
			return false
		}
	}

	return !(a+b < c || a+c < b || b+c < a)
}
