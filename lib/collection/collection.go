package collection

func Map[T any, R any](slice []T, predicate func(T) R) []R {
	newSlice := make([]R, len(slice))

	for i, it := range slice {
		newSlice[i] = predicate(it)
	}

	return newSlice
}
