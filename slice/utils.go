package slice

import (
	"strings"
)

func Map[T any, T2 any](slice []T, mapFn func(T) T2) []T2 {
	res := make([]T2, len(slice))

	for i := range slice {
		res[i] = mapFn(slice[i])
	}

	return res
}

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

func Join[T any](slice []T, separator string, stringer func(T) string) string {
	if len(slice) == 0 {
		return ""
	}

	var strArr []string = make([]string, len(slice))

	for i := range slice {
		strArr[i] = stringer(slice[i])
	}

	return strings.Join(strArr, separator)
}
