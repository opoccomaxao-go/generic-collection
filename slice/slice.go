package slice

type Slice[T any] []T

func (s Slice[T]) Slice() []T {
	return s[:]
}

//nolint:nonamedreturns // shortest possible
func (Slice[T]) Zero() (_ T) {
	return
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s *Slice[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Slice[T]) Pop() T {
	var res T

	if last := len(*s) - 1; last >= 0 {
		res = (*s)[last]
		(*s)[last] = s.Zero()
		*s = (*s)[:last]
	}

	return res
}

func (s Slice[T]) Filter(predicate func(T) bool) Slice[T] {
	return Filter(s, predicate)
}

func (s Slice[T]) Join(separator string, stringer func(T) string) string {
	return Join(s, separator, stringer)
}

// ForEach - chainable ForEach call.
func (s Slice[T]) ForEach(callbackFn func(T)) Slice[T] {
	ForEach(s, callbackFn)

	return s
}

// Range executes predicate with each element of slice until false returned.
func (s Slice[T]) Range(predicate func(T) bool) Slice[T] {
	for _, value := range s {
		if !predicate(value) {
			break
		}
	}

	return s
}

// RangeReverse executes predicate with each element of slice in reverse order until false returned.
func (s Slice[T]) RangeReverse(predicate func(T) bool) Slice[T] {
	for i := len(s) - 1; i >= 0; i-- {
		if !predicate(s[i]) {
			break
		}
	}

	return s
}

func (s Slice[T]) Chunk(size int) [][]T {
	return Chunk(s, size)
}

func WrapSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}
