package slice

import (
	"strings"

	"github.com/opoccomaxao-go/generic-collection/gmath"
	"github.com/opoccomaxao-go/generic-collection/set"
	"github.com/opoccomaxao-go/generic-collection/slice/filter"
	"github.com/opoccomaxao-go/generic-collection/slice/mapper"
)

// Map creates a new array populated with the results of calling
// a provided function on every element in the calling array.
// analogue of js Array.map.
func Map[T any, T2 any](slice []T, mapFn func(T) T2) []T2 {
	res := make([]T2, len(slice))

	for i := range slice {
		res[i] = mapFn(slice[i])
	}

	return res
}

// Filter creates a shallow copy of a portion of a given array,
// filtered down to just the elements from the given array that
// pass the test implemented by the provided function.
// analogue of js Array.filter.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(slice))

	for i := range slice {
		if predicate(slice[i]) {
			res = append(res, slice[i])
		}
	}

	return res
}

func IndexOf[T comparable](slice []T, element T) int {
	for i := range slice {
		if slice[i] == element {
			return i
		}
	}

	return -1
}

// Join - analogue of js Array.join.
func Join[T any](slice []T, separator string, stringer func(T) string) string {
	if len(slice) == 0 {
		return ""
	}

	strArr := make([]string, len(slice))

	for i := range slice {
		strArr[i] = stringer(slice[i])
	}

	return strings.Join(strArr, separator)
}

// Reduce - analogue of js Array.reduce.
func Reduce[T any, R any](
	input []T,
	reducer func(prev R, cur T) R,
	initial R,
) R {
	res := initial

	for i := range input {
		res = reducer(res, input[i])
	}

	return res
}

// UniqueUnordered - remove duplicates from slice without initial order.
func UniqueUnordered[T comparable](slice []T) []T {
	set := set.New[T]()

	for _, v := range slice {
		set.Add(v)
	}

	return set.Slice()
}

// NilInit returns new empty slice if slice is nil.
// Best use: init empty array before json.Marshal to marshal array instead of null.
func NilInit[T any](slice []T) []T {
	if slice == nil {
		return []T{}
	}

	return slice
}

// FindIndex - returns the index of the first element in an array that satisfies the provided condition.
// analogue of js.Array.findIndex.
func FindIndex[T any](slice []T, condition func(T) bool) int {
	for i, v := range slice {
		if condition(v) {
			return i
		}
	}

	return -1
}

// Find - returns the first element in the provided array that satisfies the provided testing function. If no values satisfy the testing function, zero of type is returned.
// analogue of js.Array.find.
func Find[T any](slice []T, condition func(T) bool) T {
	for _, v := range slice {
		if condition(v) {
			return v
		}
	}

	var zero T

	return zero
}

// ForEach - calls a provided callbackFn function once for each element in an array in ascending index order.
// analogue of js.Array.forEach.
func ForEach[T any](slice []T, callbackFn func(T)) {
	for _, v := range slice {
		callbackFn(v)
	}
}

func FirstNonEmpty[T comparable](values ...T) T {
	return Find(values, filter.NonEmpty[T])
}

// Nullify - assign zero to each element and reslice to zero len.
func Nullify[T any](slice []T) []T {
	var zero T

	for i := range slice {
		slice[i] = zero
	}

	return slice[0:0]
}

// Chunk split slice into chunks of size chunkSize.
func Chunk[T any](slice []T, chunkSize int) [][]T {
	total := len(slice)
	res := make([][]T, 0, total/chunkSize+1)

	for start := 0; start < total; start += chunkSize {
		end := start + chunkSize
		if end > total {
			end = total
		}

		res = append(res, slice[start:end])
	}

	return res
}

// Concat - is used to merge two or more slices.
// This method does not change the existing slices, but instead returns a new slice.
func Concat[T any](slices ...[]T) []T {
	totalLen := gmath.Sum(Map(slices, mapper.Len[T])...)
	res := make([]T, 0, totalLen)

	for _, slice := range slices {
		res = append(res, slice...)
	}

	return res
}
