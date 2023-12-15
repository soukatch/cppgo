package algorithm

import (
	"cmp"
	"gocpp/utility"
	"slices"
	"sort"
)

// Returns an iterator to the beginning of the sequence represented by c.
func Begin[T any](c []T) int {
	return 0
}

// Returns an iterator one past the end of the sequence represented by c.
func End[T any](c []T) int {
	return len(c)
}

// Checks if unary predicate p returns true for all elements in the range
// r[first, last).
func AllOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIfNot(r, first, last, p) == last
}

// Checks if unary predicate p returns true for at least one element in the
// range r[first, last).
func AnyOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIf(r, first, last, p) != last
}

// Checks if unary predicate p returns true for no elements in the range
// r[first, last).
func NoneOf[T any](r []T, first, last int, p func(T) bool) bool {
	return FindIf(r, first, last, p) == last
}

// Searches for an element equal to value using operator==.
func Find[T comparable](r []T, first, last int, value T) int {
	return slices.Index(r[first:last], value)
}

// Searches for an element for which predicate p returns true.
func FindIf[T any](r []T, first, last int, p func(T) bool) int {
	return slices.IndexFunc(r[first:last], p)
}

// Searches for an element for which predicate q returns false.
func FindIfNot[T any](r []T, first, last int, q func(T) bool) int {
	return slices.IndexFunc(r[first:last], func(x T) bool { return !q(x) })
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that element >= value, or last if no such element is found by invoking
// slices.BinarySearch().
func LowerBound[T cmp.Ordered](r []T, first, last int, value T) int {
	first, _ = slices.BinarySearch(r[first:last], value)
	return first
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that comp(element, value) is false, or last if no such element is found
// by invoking slices.BinarySearch().
func LowerBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) int) int {
	first, _ = slices.BinarySearchFunc(r[first:last], value, comp)
	return first
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that value < element, or last if no such element is found, by invoking
// sort.Search().
func UpperBound[T cmp.Ordered](r []T, first, last int, value T) int {
	first = sort.Search(len(r[first:last]), func(i int) bool {
		return r[i] > value
	})
	return first
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that comp(element, value) is true, or last if no such element is found,
// by invoking sort.Search().
func UpperBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) int) int {
	first = sort.Search(len(r[first:last]), func(i int) bool {
		return comp(r[i], value) == 1
	})
	return first
}

// Searches the range [first, last) for two consecutive equal elements. Returns
// an iterator to the first of the first pair of identical elements if found,
// that is, the first iterator it such that r[it] == r[it + 1]; last otherwise.
func AdjacentFind[T comparable](r []T, first, last int) int {
	if first == last {
		return last
	}

	for next := first + 1; next != last; {
		if r[first] == r[next] {
			return first
		}
		next++
		first++
	}

	return last
}

// Searches the range r[first, last) for two consecutive equal elements. Returns
// an iterator to the first of the first pair of identical elements if found,
// that is, the first iterator it such that p(*it, *(it + 1)) != false; last
// otherwise.
func AdjacentFindFunc[T any](r []T, first, last int, p func(T, T) bool) int {
	if first == last {
		return last
	}

	for next := first + 1; next != last; {
		if p(r[first], r[next]) {
			return first
		}
		next++
		first++
	}

	return last
}

func Search[T comparable](r1, r2 []T, first, last, s_first, s_last int) int {
	for {
		it := first
		for s_it := s_first; ; {
			if s_it == s_last {
				return first
			}
			if it == last {
				return last
			}
			if r1[it] != r2[s_it] {
				break
			}
			it++
			s_it++
		}
		first++
	}
}

func SearchFunc[T any](r1, r2 []T, first, last, s_first, s_last int, p func(T, T) bool) int {
	for {
		it := first
		for s_it := s_first; ; {
			if s_it == s_last {
				return first
			}
			if it == last {
				return last
			}
			if !p(r1[it], r2[s_it]) {
				break
			}
			it++
			s_it++
		}
		first++
	}
}

func SearchN[T comparable](r []T, first, last, count int, value T) int {
	if count <= 0 {
		return first
	}

	for ; first != last; first++ {
		if r[first] != value {
			continue
		}

		candidate := first

		for cur_count := 1; ; cur_count++ {
			if cur_count >= count {
				return candidate
			}

			first++
			if first == last {
				return last
			}

			if r[first] != value {
				break
			}
		}
	}
	return last
}

func SearchNFunc[T any](r []T, first, last, count int, value T, p func(T, T) bool) int {
	if count <= 0 {
		return first
	}

	for ; first != last; first++ {
		if !p(r[first], value) {
			continue
		}

		candidate := first

		for cur_count := 1; ; cur_count++ {
			if cur_count >= count {
				return candidate
			}

			if first++; first == last {
				return last
			}

			if !p(r[first], value) {
				break
			}
		}
	}
	return last
}

func Mismatch[T comparable](r1, r2 []T, first1, last1, first2 int) utility.Pair[int, int] {
	for first1 != last1 && r1[first1] == r1[first2] {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}

func MismatchFunc[T any](r1, r2 []T, first1, last1, first2 int, p func(T, T) bool) utility.Pair[int, int] {
	for first1 != last1 && p(r1[first1], r2[first2]) {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}

func Mismatch2[T comparable](r1, r2 []T, first1, last1, first2, last2 int) utility.Pair[int, int] {
	for first1 != last1 && first2 != last2 && r1[first1] == r1[first2] {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}

func MismatchFunc2[T any](r1, r2 []T, first1, last1, first2, last2 int, p func(T, T) bool) utility.Pair[int, int] {
	for first1 != last1 && first2 != last2 && p(r1[first1], r2[first2]) {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}
