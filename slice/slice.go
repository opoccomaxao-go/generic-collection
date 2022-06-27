package slice

type Slice[T any] []T

func (s Slice[T]) Slice() []T {
	return s[:]
}

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

	last := len(*s) - 1
	if last >= 0 {
		res = (*s)[last]
		(*s)[last] = s.Zero()
		*s = (*s)[:last]
	}

	return res
}

func (s Slice[T]) Filter(predicate func(T) bool) Slice[T] {
	return Filter(s, predicate)
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

func WrapSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}
