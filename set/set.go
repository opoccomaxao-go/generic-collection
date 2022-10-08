package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return Set[T]{}
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]

	return ok
}

func (s Set[T]) Delete(value T) {
	delete(s, value)
}

func (s Set[T]) Slice() []T {
	res := make([]T, 0, len(s))

	for key := range s {
		res = append(res, key)
	}

	return res
}

func (s Set[T]) SliceSorted(sortFns ...func([]T)) []T {
	res := s.Slice()

	for _, sortFn := range sortFns {
		sortFn(res)
	}

	return res
}
