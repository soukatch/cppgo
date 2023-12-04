package algorithm

import (
	"cmp"
	"slices"
	"sort"
)

// Returns an iterator to the beginning of the sequence represented by c.
func Begin[T any](c []T) *T {
	return &c[0]
}

// Returns an iterator one past the end of the sequence represented by c.
func End[T any](c []T) *T {
	return &c[len(c)]
}

// Checks if unary predicate p returns true for all elements in the range
// [first, last).
func AllOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIfNot(r, first, last, p) == &r[last]
}

// Checks if unary predicate p returns true for at least one element in the
// range [first, last).
func AnyOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIf(r, first, last, p) != &r[last]
}

// Checks if unary predicate p returns true for no elements in the range
// [first, last).
func NoneOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIf(r, first, last, p) == &r[last]
}

// Searches for an element equal to value using operator==.
func Find[T comparable](r []T, first, last int, value T) *T {
	return &r[slices.Index(r[first:last], value)]
}

// Searches for an element for which predicate p returns true.
func FindIf[T any](r []T, first, last int, p func(T) bool) *T {
	return &r[slices.IndexFunc(r[first:last], p)]
}

// Searches for an element for which predicate q returns false.
func FindIfNot[T any](r []T, first, last int, q func(T) bool) *T {
	return &r[slices.IndexFunc(r[first:last], func(x T) bool { return !q(x) })]
}

// Returns an iterator pointing to the first element in the range [first, last)
// such that element >= value, or last if no such element is found by invoking
// slices.BinarySearch().
func LowerBound[T cmp.Ordered](r []T, first, last int, value T) *T {
	first, _ = slices.BinarySearch(r[first:last], value)
	return &r[first]
}

// Returns an iterator pointing to the first element in the range [first, last)
// such that comp(element, value) is false, or last if no such element is found
// by invoking slices.BinarySearch().
func LowerBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) int) *T {
	first, _ = slices.BinarySearchFunc(r[first:last], value, comp)
	return &r[first]
}

// Returns an iterator pointing to the first element in the range [first, last)
// such that value < element, or last if no such element is found, by invoking
// sort.Search().
func UpperBound[T cmp.Ordered](r []T, first, last int, value T) *T {
	first = sort.Search(len(r[first:last]), func(i int) bool {
		return r[i] > value
	})
	return &r[first]
}

// Returns an iterator pointing to the first element in the range [first, last)
// such that comp(element, value) is true, or last if no such element is found,
// by invoking sort.Search().
func UpperBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) int) *T {
	first = sort.Search(len(r[first:last]), func(i int) bool {
		return comp(r[i], value) == 1
	})
	return &r[first]
}

// Searches the range [first, last) for two consecutive equal elements. Returns
// an iterator to the first of the first pair of identical elements if found,
// that is, the first iterator it such that *it == *(it + 1); last otherwise.
func AdjacentFind[T comparable](r []T, first, last int) *T {
	if first == last {
		return &r[last]
	}

	for next := first + 1; next != last; {
		if r[first] == r[next] {
			return &r[first]
		}
		next++
		first++
	}

	return &r[last]
}

// Searches the range [first, last) for two consecutive equal elements. Returns
// an iterator to the first of the first pair of identical elements if found,
// that is, the first iterator it such that p(*it, *(it + 1)) != false; last
// otherwise.
func AdjacentFindFunc[T any](r []T, first, last int, p func(T, T) bool) *T {
	if first == last {
		return &r[last]
	}

	for next := first + 1; next != last; {
		if p(r[first], r[next]) {
			return &r[first]
		}
		next++
		first++
	}

	return &r[last]
}
