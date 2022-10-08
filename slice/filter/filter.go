package filter

func NonEmpty[T comparable](item T) bool {
	var zero T

	return zero != item
}
