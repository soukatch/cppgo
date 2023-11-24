package algorithm

import (
	"cmp"
	"slices"
	"sort"
)

func AllOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIfNot(r, first, last, p) == &r[last]
}

func AnyOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIf(r, first, last, p) != &r[last]
}

func NoneOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIf(r, first, last, p) == &r[last]
}

func Find[T comparable](r []T, first, last int, value T) *T {
	return &r[slices.Index(r[first:last], value)]
}

func FindIf[T any](r []T, first, last int, p func(T) bool) *T {
	return &r[slices.IndexFunc(r[first:last], p)]
}

func FindIfNot[T any](r []T, first, last int, q func(T) bool) *T {
	return &r[slices.IndexFunc(r[first:last], func(x T) bool { return !q(x) })]
}

func LowerBound[T cmp.Ordered](r []T, first, last int, value T) *T {
	first, _ = slices.BinarySearch(r[first:last], value)
	return &r[first]
}

func LowerBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) int) *T {
	first, _ = slices.BinarySearchFunc(r[first:last], value, comp)
	return &r[first]
}

func UpperBound[T cmp.Ordered](r []T, first, last int, value T) *T {
	first = sort.Search(len(r[first:last]), func(i int) bool {
		return r[i] > value
	})
	return &r[first]
}

func UpperBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) int) *T {
	first = sort.Search(len(r[first:last]), func(i int) bool {
		return comp(r[i], value) == 1
	})
	return &r[first]
}
