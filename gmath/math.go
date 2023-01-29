// package gmath - generic math.
//
// Contains common math functions with generic signature.
package gmath

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](first T, others ...T) T {
	res := first

	for _, v := range others {
		if v > res {
			res = v
		}
	}

	return res
}

func Min[T constraints.Ordered](first T, others ...T) T {
	res := first

	for _, v := range others {
		if v < res {
			res = v
		}
	}

	return res
}

func Abs[T constraints.Integer | constraints.Float](value T) T {
	if value < 0 {
		return -value
	}

	return value
}

func Sum[T constraints.Integer | constraints.Float](values ...T) T {
	var res T = 0

	for _, value := range values {
		res += value
	}

	return res
}
