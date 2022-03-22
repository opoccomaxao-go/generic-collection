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

func WrapSlice[T any](slice []T) Slice[T] {
	return Slice[T](slice)
}
