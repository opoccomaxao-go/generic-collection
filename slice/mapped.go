package slice

type MappedSlice[T any, T2 any] struct {
	Slice[T]
}

func (s *MappedSlice[T, T2]) Map(mapFn func(T) T2) Slice[T2] {
	return Map(s.Slice, mapFn)
}

func WrapMapped[T any, T2 any](slice []T) MappedSlice[T, T2] {
	return MappedSlice[T, T2]{
		Slice: slice,
	}
}
