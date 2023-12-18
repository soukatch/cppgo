package algorithm

import (
	"cmp"
	"gocpp/utility"
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

// Searches for an element equal to value (using operator==).
func Find[T comparable](r []T, first, last int, value T) int {
	for ; first != last; first++ {
		if r[first] == value {
			return first
		}
	}
	return last
}

// Searches for an element for which predicate p returns true.
func FindIf[T any](r []T, first, last int, p func(T) bool) int {
	for ; first != last; first++ {
		if p(r[first]) {
			return first
		}
	}
	return last
}

// Searches for an element for which predicate q returns false.
func FindIfNot[T any](r []T, first, last int, q func(T) bool) int {
	for ; first != last; first++ {
		if !q(r[first]) {
			return first
		}
	}
	return last
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that element >= value, or last if no such element is found.
func LowerBound[T cmp.Ordered](r []T, first, last int, value T) int {
	for first < last-1 {
		if it := first + (last-first)/2; r[it] < value {
			first = it
		} else {
			last = it
		}
	}
	return last
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that comp(element, value) is false, or last if no such element is found.
func LowerBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) bool) int {
	for first < last-1 {
		if it := first + (last-first)/2; comp(r[it], value) {
			first = it
		} else {
			last = it
		}
	}
	return last
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that value < element, or last if no such element is found.
func UpperBound[T cmp.Ordered](r []T, first, last int, value T) int {
	for first < last-1 {
		if it := first + (last-first)/2; r[it] <= value {
			first = it
		} else {
			last = it
		}
	}
	return last
}

// Returns an iterator pointing to the first element in the range r[first, last)
// such that comp(element, value) is true, or last if no such element is found.
func UpperBoundFunc[T any](r []T, first, last int, value T, comp func(T, T) bool) int {
	for first < last-1 {
		if it := first + (last-first)/2; !comp(value, r[it]) {
			first = it
		} else {
			last = it
		}
	}
	return last
}

// Searches the range r[first, last) for two consecutive equal elements. Returns
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

// Returns the number of elements in the range r[first, last) satisfying
// specific criteria. Counts the elements that are equal to value
// (using operator==).
func Count[T comparable](r []T, first, last int, value T) int {
	ret := int(0)
	for ; first != last; first++ {
		if r[first] == value {
			ret++
		}
	}
	return ret
}

// Returns the number of elements in the range r[first, last) satisfying specific
// criteria. Counts elements for which predicate p returns true.
func CountIf[T any](r []T, first, last int, p func(T) bool) int {
	ret := int(0)
	for ; first != last; first++ {
		if p(r[first]) {
			ret++
		}
	}
	return ret
}

// Returns the first mismatching pair of elements from two ranges: one defined
// by r1[first1, last1) and another defined by r2[first2, first2 + last1 - first1).
// Elements are compared using operator==.
func Mismatch[T comparable](r1, r2 []T, first1, last1, first2 int) utility.Pair[int, int] {
	for first1 != last1 && r1[first1] == r1[first2] {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}

// Returns the first mismatching pair of elements from two ranges: one defined
// by r1[first1, last1) and another defined by r2[first2, first2 + last1 - first1).
// Elements are compared using the given binary predicate p.
func MismatchFunc[T any](r1, r2 []T, first1, last1, first2 int, p func(T, T) bool) utility.Pair[int, int] {
	for first1 != last1 && p(r1[first1], r2[first2]) {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}

// Returns the first mismatching pair of elements from two ranges: one defined
// by r1[first1, last1) and another defined by r2[first2, last2). Elements are
// compared using operator==.
func Mismatch2[T comparable](r1, r2 []T, first1, last1, first2, last2 int) utility.Pair[int, int] {
	for first1 != last1 && first2 != last2 && r1[first1] == r1[first2] {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}

// Returns the first mismatching pair of elements from two ranges: one defined
// by r1[first1, last1) and another defined by r2[first2, last2). Elements are
// compared using the given binary predicate p.
func MismatchFunc2[T any](r1, r2 []T, first1, last1, first2, last2 int, p func(T, T) bool) utility.Pair[int, int] {
	for first1 != last1 && first2 != last2 && p(r1[first1], r2[first2]) {
		first1++
		first2++
	}

	return utility.MakePair(first1, first2)
}

// Returns true if the range r1[first1, last1) is equal to the range r2[first2,
// first2 + (last1 - first1)), and false otherwise. Two ranges are considered
// equal if they have the same number of elements and, for every iterator i in
// the range [first1, last1), *i equals *(first2 + (i - first1)). Uses
// operator== to determine if two elements are equal.
func Equal[T comparable](r1, r2 []T, first1, last1, first2 int) bool {
	for first1 != last1 {
		if r1[first1] != r2[first2] {
			return false
		}
		first1++
		first2++
	}

	return true
}

// Returns true if the range r1[first1, last1) is equal to the range r2[first2,
// first2 + (last1 - first1)), and false otherwise. Two ranges are considered
// equal if they have the same number of elements and, for every iterator i in
// the range [first1, last1), *i equals *(first2 + (i - first1)). Uses given
// binary predicate p to determine if two elements are equal.
func EqualFunc[T any](r1, r2 []T, first1, last1, first2 int, p func(T, T) bool) bool {
	for first1 != last1 {
		if !p(r1[first1], r2[first2]) {
			return false
		}
		first1++
		first2++
	}

	return true
}

// Returns true if the range r1[first1, last1) is equal to the range r2[first2,
// last2), and false otherwise. Two ranges are considered equal if they have
// the same number of elements and, for every iterator i in the range [first1,
// last1), *i equals *(first2 + (i - first1)). Uses operator== to determine if
// two elements are equal.
func Equal2[T comparable](r1, r2 []T, first1, last1, first2, last2 int) bool {
	for first1 != last1 && first2 != last2 {
		if r1[first1] != r2[first2] {
			return false
		}
		first1++
		first2++
	}

	return true
}

// Returns true if the range r1[first1, last1) is equal to the range r2[first2,
// last2), and false otherwise. Two ranges are considered equal if they have
// the same number of elements and, for every iterator i in the range [first1,
// last1), *i equals *(first2 + (i - first1)). Uses given binary predicate p to
// determine if two elements are equal.
func EqualFunc2[T any](r1, r2 []T, first1, last1, first2, last2 int, p func(T, T) bool) bool {
	for first1 != last1 && first2 != last2 {
		if !p(r1[first1], r2[first2]) {
			return false
		}
		first1++
		first2++
	}

	return true
}

// Searches for the first occurrence of the sequence of elements r2[s_first,
// s_last) in the range r1[first, last). Elements are compared using operator==.
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

// Searches for the first occurrence of the sequence of elements r2[s_first,
// s_last) in the range r1[first, last). Elements are compared using the given
// binary predicate p.
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

// Searches the range r[first, last) for the first sequence of count identical
// elements, each equal to the given value. Elements are compared using
// operator==.
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

// Searches the range r[first, last) for the first sequence of count identical
// elements, each equal to the given value. Elements are compared using the
// given binary predicate p.
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

// Copies the elements in the range, defined by r[first, last), to another range
// beginning at r[d_first] (copy destination range). Copies all elements in the
// range r[first, last) starting from first and proceeding to last. If d_first
// is in r[first, last), the behavior is undefined.
func Copy[T any](r1, r2 []T, first, last, d_first int) int {
	for first != last {
		r2[d_first] = r1[first]
		first++
		d_first++
	}

	return d_first
}

// Copies the elements in the range, defined by r[first, last), to another range
// beginning at r[d_first] (copy destination range). Only copies the elements
// for which the predicate pred returns true. This copy algorithm is stable: the
// relative order of the elements that are copied is preserved. If [first, last)
// and the copy destination range overlaps, the behavior is undefined.
func CopyIf[T any](r1, r2 []T, first, last, d_first int, pred func(T) bool) int {
	for first != last {
		if pred(r1[first]) {
			r2[d_first] = r1[first]
			d_first++
		}
		first++
	}

	return d_first
}

// Copies exactly count values from the range beginning at r1[first] to the range
// beginning at r2[result]. Formally, for each integer 0 ≤ i < count, performs
// *(result + i) = *(first + i). Overlap of ranges is formally permitted, but
// leads to unpredictable ordering of the results.
func CopyN[T any](r1, r2 []T, first, count, result int) int {
	if count > 0 {
		r2[result] = r1[first]
		first++
		result++
		for i := 1; i != count; i++ {
			r2[result] = r1[first]
			first++
			result++
		}
	}

	return result
}

// Copies the elements from the range r1[first, last) to another range ending at
// r2[d_last]. The elements are copied in reverse order (the last element is
// copied first), but their relative order is preserved. The behavior is
// undefined if d_last is within (first, last).
func CopyBackward[T any](r1, r2 []T, first, last, d_last int) int {
	for first != last {
		d_last--
		last--
		r2[d_last] = r1[last]
	}
	return d_last
}

// Moves the elements in the range r1[first, last), to another range beginning
// at r2[d_first], starting from r1[first] and proceeding to r1[last - 1].
func Move[T any](r1, r2 []T, first, last, d_first int) int {
	return Copy(r1, r2, first, last, d_first)
}

// Moves the elements from the range r1[first, last), to another range ending at
// r2[d_last]. The elements are moved in reverse order (the last element is
// moved first), but their relative order is preserved.
func MoveBackward[T any](r1, r2 []T, first, last, d_first int) int {
	return CopyBackward(r1, r2, first, last, d_first)
}

// Swaps the values a and b.
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// Exchanges elements between range r1[first1, last1) and another range starting
// at r2[first2]. Precondition: the two ranges f1[first1, last1) and r2[first2,
// last2) do not overlap,
// where r2[last2] = r2[Next(first2, distance(first1, last1))].
func SwapRanges[T any](r1, r2 []T, first1, last1, first2 int) int {
	for first1 != last1 {
		IterSwap(&r1[first1], &r2[first2])
		first1++
		first2++
	}
	return first2
}

// Swaps the values of the elements the given iterators are pointing to.
func IterSwap[T any](a, b *T) {
	*a, *b = *b, *a
}

// Applies the given function to a range and stores the result in another range,
// keeping the original elements order and beginning at r2[d_first]. The unary
// operation unary_op is applied to the range defined by r1[first1, last1).
func Transform[T1, T2 any](r1 []T1, r2 []T2, first1, last1, d_first int, unary_op func(T1) T2) int {
	for first1 != last1 {
		r2[d_first] = unary_op(r1[first1])
		first1++
		d_first++
	}
	return d_first
}

// Applies the given function to a range and stores the result in another range,
// keeping the original elements order and beginning at r2[d_first]. The binary
// operation binary_op is applied to pairs of elements from two ranges: one
// defined by [first1, last1) and the other beginning at first2.
func Transform2[T1, T2, T3 any](r1 []T1, r2 []T2, r3 []T3, first1, last1, first2, d_first int, binary_op func(T1, T2) T3) int {
	for first1 != last1 {
		r3[d_first] = binary_op(r1[first1], r2[first2])
		first1++
		first2++
		d_first++
	}
	return d_first
}

// Replaces all elements satisfying specific criteria with new_value in the
// range r[first, last). Replaces all elements that are equal to old_value
// (using operator==).
func Replace[T comparable](r []T, first, last int, old_value, new_value T) {
	for ; first != last; first++ {
		if r[first] == old_value {
			r[first] = new_value
		}
	}
}

// Replaces all elements satisfying specific criteria with new_value in the
// range r[first, last). Replaces all elements for which predicate p returns
// true.
func ReplaceIf[T any](r []T, first, last int, p func(T) bool, new_value T) {
	for ; first != last; first++ {
		if p(r[first]) {
			r[first] = new_value
		}
	}
}

// Copies the elements from the range r1[first, last) to another range beginning
// at r2[d_first], while replacing all elements satisfying specific criteria
// with new_value. If the source and destination ranges overlap, the behavior is
// undefined. Replaces all elements that are equal to old_value
// (using operator==).
func ReplaceCopy[T comparable](r1, r2 []T, first, last, d_first int, old_value, new_value T) int {
	for ; first != last; first++ {
		r2[d_first] = r1[first]
		if r2[d_first] == old_value {
			r2[d_first] = new_value
		}
		d_first++
	}
	return d_first
}

// Copies the elements from the range r1[first, last) to another range beginning
// at r2[d_first], while replacing all elements satisfying specific criteria
// with new_value. If the source and destination ranges overlap, the behavior is
// undefined. Replaces all elements for which predicate p returns true.
func ReplaceCopyIf[T any](r1, r2 []T, first, last, d_first int, p func(T) bool, new_value T) int {
	for ; first != last; first++ {
		r2[d_first] = r1[first]
		if p(r2[d_first]) {
			r2[d_first] = new_value
		}
		d_first++
	}
	return d_first
}

// Assigns the given value to the elements in the range r[first, last).
func Fill[T any](r []T, first, last int, value T) {
	for ; first != last; first++ {
		r[first] = value
	}
}

// Assigns the given value to the first count elements in the range beginning at
// first if count > 0. Does nothing otherwise.
func FillN[T any](r []T, first, count int, value T) int {
	for i := 0; i < count; i++ {
		r[first] = value
		first++
	}
	return first
}

// Assigns each element in range r[first, last) a value generated by the given
// function object g.
func Generate[T any](r []T, first, last int, g func() T) {
	for ; first != last; first++ {
		r[first] = g()
	}
}

// Assigns values, generated by given function object g, to the first count
// elements in the range beginning at first, if count > 0. Does nothing
// otherwise.
func GenerateN[T any](r []T, first, count int, g func() T) int {
	for i := 0; i < count; i++ {
		r[first] = g()
		first++
	}

	return first
}

// Removes all elements satisfying specific criteria from the range [first,
// last) and returns a past-the-end iterator for the new end of the range.
// Removes all elements that are equal to value (using operator==).
func Remove[T comparable](r []T, first, last int, value T) int {
	first = Find(r, first, last, value)
	if first != last {
		for i := first + 1; i != last; i++ {
			if r[i] == value {
				r[first] = r[i]
				first++
			}
		}
	}
	return first
}

// Removes all elements satisfying specific criteria from the range [first,
// last) and returns a past-the-end iterator for the new end of the range.
// Removes all elements for which predicate p returns true.
func RemoveIf[T any](r []T, first, last int, p func(T) bool) int {
	first = FindIf(r, first, last, p)
	if first != last {
		for i := first + 1; i != last; i++ {
			if !p(r[i]) {
				r[first] = r[i]
				first++
			}
		}
	}
	return first
}

// Copies elements from the range r1[first, last), to another range beginning at
// r2[d_first], omitting the elements which satisfy specific criteria. Ignores
// all elements that are equal to value.
func RemoveCopy[T comparable](r1, r2 []T, first, last, d_first int, value T) int {
	for ; first != last; first++ {
		if r1[first] != value {
			r2[d_first] = r1[first]
			d_first++
		}
	}
	return d_first
}

// Copies elements from the range r1[first, last), to another range beginning at
// r2[d_first], omitting the elements which satisfy specific criteria. Ignores
// all elements for which predicate p returns true.
func RemoveCopyIf[T any](r1, r2 []T, first, last, d_first int, p func(T) bool) int {
	for ; first != last; first++ {
		if !p(r1[first]) {
			r2[d_first] = r1[first]
			d_first++
		}
	}
	return d_first
}

// Eliminates all except the first element from every consecutive group of
// equivalent elements from the range r[first, last) and returns a past-the-end
// iterator for the new logical end of the range. Removing is done by shifting
// the elements in the range in such a way that elements to be erased are
// overwritten. Elements are compared using operator==. The behavior is
// undefined if it is not an equivalence relation.
func Unique[T comparable](r []T, first, last int) int {
	if first == last {
		return last
	}

	result := first
	for first++; first != last; first++ {
		if r[result] != r[first] {
			result++
			// if result != first {
			r[result] = r[first]
			// }
		}
	}
	return result + 1
}

// Eliminates all except the first element from every consecutive group of
// equivalent elements from the range r[first, last) and returns a past-the-end
// iterator for the new logical end of the range. Removing is done by shifting
// the elements in the range in such a way that elements to be erased are
// overwritten. Elements are compared using the given binary predicate p. The
// behavior is undefined if it is not an equivalence relation.
func UniqueFunc[T any](r []T, first, last int, p func(T, T) bool) int {
	if first == last {
		return last
	}

	result := first
	for first++; first != last; first++ {
		if !p(r[result], r[first]) {
			result++
			// if result != first {
			r[result] = r[first]
			// }
		}
	}

	return result + 1
}

// Copies the elements from the range r[first, last), to another range beginning
// at d_first in such a way that there are no consecutive equal elements. Only
// the first element of each group of equal elements is copied. Elements are
// compared using operator==. The behavior is undefined if it is not an
// equivalence relation.
func UniqueCopy[T comparable](r1, r2 []T, first, last, d_first int) int {
	if first == last {
		return d_first
	}

	r2[d_first] = r1[first]

	for first++; first != last; first++ {
		if !(r1[first] == r2[d_first]) {
			d_first++
			r2[d_first] = r1[first]
		}
	}

	return d_first + 1
}

// Copies the elements from the range r[first, last), to another range beginning
// at d_first in such a way that there are no consecutive equal elements. Only
// the first element of each group of equal elements is copied. Elements are
// compared using the given binary predicate p. The behavior is undefined if it
// is not an equivalence relation.
func UniqueCopyFunc[T any](r1, r2 []T, first, last, d_first int, p func(T, T) bool) int {
	if first == last {
		return d_first
	}

	r2[d_first] = r1[first]

	for first++; first != last; first++ {
		if !p(r1[first], r2[d_first]) {
			d_first++
			r2[d_first] = r1[first]
		}
	}

	return d_first + 1
}

// Reverses the order of the elements in the range r[first, last). Behaves as if
// applying IterSwap to every pair of iterators first + i and (last - i) - 1 for
// each integer i in [​0​, Distance(first, last) / 2).
func Reverse[T any](r []T, first, last int) {
	for last--; first < last; {
		IterSwap(&r[first], &r[last])
		first++
		last--
	}
}

// Given  N as Distance(first, last). Copies the elements from the range [first,
// last) to another range of  N elements beginning at d_first (destination
// range) in such a way that the elements in the destination range are in
// reverse order. Behaves as if by executing the assignment *(d_first + N - 1 -
// i) = *(first + i) once for each integer i in [​0​, N). If [first, last) and
// the destination range overlap, the behavior is undefined.
func ReverseCopy[T any](r1, r2 []T, first, last, d_first int) int {
	for ; first != last; d_first++ {
		last--
		r2[d_first] = r1[last]
	}
	return d_first
}

// Performs a left rotation on a range of elements. Specifically, Rotate swaps
// the elements in the range [first, last) in such a way that the elements in
// r[first, middle) are placed after the elements in [middle, last) while the
// orders of the elements in both ranges are preserved. If r[first, middle) or
// r[middle, last) is not a valid range, the behavior is undefined.
func Rotate[T any](r []T, first, middle, last int) int {
	if first == middle {
		return last
	}

	if middle == last {
		return first
	}

	write, next_read := first, first

	for read := middle; read != last; {
		if write == next_read {
			next_read = read
		}
		IterSwap(&r[write], &r[read])
		write++
		read++
	}

	Rotate(r, write, next_read, last)
	return write
}

// Copies the elements from the range r1[first, last), to another range
// beginning at r2[d_first] in such a way, that the element *(n_first) becomes
// the first element of the new range and *(n_first - 1) becomes the last
// element. The behavior is undefined if either [first, n_first) or [n_first,
// last) is not a valid range, or the source and destination ranges overlap.
func RotateCopy[T any](r1, r2 []T, first, n_first, last, d_first int) int {
	return Copy(r1, r2, first, n_first, Copy(r1, r2, n_first, last, d_first))
}

// Shifts the elements towards the beginning of the range. If n == 0 ||
// n >= last - first, there are no effects. If n < 0, the behavior is
// undefined. Otherwise, for every integer i in [​0​, last - first - n), moves
// the element originally at position first + n + i to position first + i.
// The moves are performed in increasing order of i starting from ​0​.
func ShiftLeft[T any](r []T, first, last, n int) int {
	if n == 0 {
		return last
	}

	if n >= last-first {
		return first
	}

	return Move(r, r, first+n, last, first)
}

// Shifts the elements towards the end of the range. If n == 0 || n >= last -
// first, there are no effects. If n < 0, the behavior is undefined. Otherwise,
// for every integer i in [​0​, last - first - n), moves the element originally
// at position first + i to position first + n + i.
func ShiftRight[T any](r []T, first, last, n int) int {
	if n == 0 {
		return last
	}

	if n >= last-first {
		return first
	}

	for i := last - 1; i >= first+n; i-- {
		r[i] = r[i-n]
	}

	return last
}

// Returns true if all elements in the range r[first, last) that satisfy the
// predicate p appear before all elements that do not. Also returns true if
// r[first, last) is empty.
func IsPartitioned[T any](r []T, first, last int, p func(T) bool) bool {
	for ; first != last; first++ {
		if !p(r[first]) {
			break
		}
	}

	for ; first != last; first++ {
		if p(r[first]) {
			return false
		}
	}

	return true
}

// Reorders the elements in the range r[first, last) in such a way that all
// elements for which the predicate p returns true precede the elements for
// which predicate p returns false. Relative order of the elements is not
// preserved. Returns iterator to the first element of the second group.
func Partition[T any](r []T, first, last int, p func(T) bool) int {
	first = FindIfNot(r, first, last, p)
	if first == last {
		return first
	}

	for i := first + 1; i != last; i++ {
		if p(r[i]) {
			IterSwap(&r[i], &r[first])
			first++
		}
	}

	return first
}

// Copies the elements from the range r1[first, last) to two different ranges
// depending on the value returned by the predicate p. The elements that satisfy
// the predicate p are copied to the range beginning at r2[d_first_true]. The
// rest of the elements are copied to the range beginning at r2[d_first_false].
// The behavior is undefined if the input range overlaps either of the output
// ranges.
func PartitionCopy[T any](r1, r2 []T, first, last, d_first_true, d_first_false int, p func(T) bool) utility.Pair[int, int] {
	for ; first != last; first++ {
		if p(r1[first]) {
			r2[d_first_true] = r1[first]
			d_first_true++
		} else {
			r2[d_first_false] = r1[first]
			d_first_false++
		}
	}

	return utility.MakePair(d_first_true, d_first_false)
}

// Examines the partitioned (as if by Partition) range r[first, last) and
// locates the end of the first partition, that is, the first element that does
// not satisfy p or last if all elements satisfy p.
func PartitionPoint[T any](r []T, first, last int, p func(T) bool) int {
	for length := last - first; 0 < length; {
		half := length / 2
		middle := first + half
		if p(r[middle]) {
			first = middle + 1
			length -= half + 1
		} else {
			length = half
		}
	}

	return first
}
