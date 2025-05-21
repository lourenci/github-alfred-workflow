package collection

import "slices"

func Map[T any, R any](slice []T, predicate func(T) R) []R {
	newSlice := make([]R, len(slice))

	for i, it := range slice {
		newSlice[i] = predicate(it)
	}

	return newSlice
}

func Dedup[T comparable](slice []T) []T {
	deduped := []T{}

	for _, it := range slice {
		if slices.Contains(deduped, it) {
			continue
		}

		deduped = append(deduped, it)
	}

	return deduped
}
