package set

func FromSlice[T comparable](slice []T) Set[T] {
	res := Set[T]{}

	for _, i := range slice {
		res[i] = struct{}{}
	}

	return res
}

func FromValues[T comparable](values ...T) Set[T] {
	return FromSlice(values)
}
