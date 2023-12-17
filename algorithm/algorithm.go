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

// Searches for an element equal to value using operator==.
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

func Count[T comparable](r []T, first, last int, value T) int {
	ret := int(0)
	for ; first != last; first++ {
		if r[first] == value {
			ret++
		}
	}
	return ret
}

func CountIf[T any](r []T, first, last int, p func(T) bool) int {
	ret := int(0)
	for ; first != last; first++ {
		if p(r[first]) {
			ret++
		}
	}
	return ret
}

func Copy[T any](r1, r2 []T, first, last, d_first int) int {
	for first != last {
		r2[d_first] = r1[first]
		first++
		d_first++
	}

	return d_first
}

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

func CopyBackward[T any](r1, r2 []T, first, last, d_last int) int {
	for first != last {
		d_last--
		last--
		r2[d_last] = r1[last]
	}
	return d_last
}

func Move[T any](r1, r2 []T, first, last, d_first int) int {
	return Copy(r1, r2, first, last, d_first)
}

func MoveBackward[T any](r1, r2 []T, first, last, d_first int) int {
	return CopyBackward(r1, r2, first, last, d_first)
}

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func SwapRanges[T any](r1, r2 []T, first1, last1, first2 int) int {
	for first1 != last1 {
		IterSwap(&r1[first1], &r2[first2])
		first1++
		first2++
	}
	return first2
}

func IterSwap[T any](a, b *T) {
	*a, *b = *b, *a
}

func Transform[T1, T2 any](r1 []T1, r2 []T2, first1, last1, d_first int, unary_op func(T1) T2) int {
	for first1 != last1 {
		r2[d_first] = unary_op(r1[first1])
		first1++
		d_first++
	}
	return d_first
}

func Transform2[T1, T2, T3 any](r1 []T1, r2 []T2, r3 []T3, first1, last1, first2, d_first int, binary_op func(T1, T2) T3) int {
	for first1 != last1 {
		r3[d_first] = binary_op(r1[first1], r2[first2])
		first1++
		first2++
		d_first++
	}
	return d_first
}

func Replace[T comparable](r []T, first, last int, old_value, new_value T) {
	for ; first != last; first++ {
		if r[first] == old_value {
			r[first] = new_value
		}
	}
}

func ReplaceIf[T any](r []T, first, last int, p func(T) bool, new_value T) {
	for ; first != last; first++ {
		if p(r[first]) {
			r[first] = new_value
		}
	}
}

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

func Fill[T any](r []T, first, last int, value T) {
	for ; first != last; first++ {
		r[first] = value
	}
}

func FillN[T any](r []T, first, count int, value T) int {
	for i := 0; i < count; i++ {
		r[first] = value
		first++
	}
	return first
}

func Generate[T any](r []T, first, last int, g func() T) {
	for ; first != last; first++ {
		r[first] = g()
	}
}

func GenerateN[T any](r []T, first, count int, g func() T) int {
	for i := 0; i < count; i++ {
		r[first] = g()
		first++
	}

	return first
}

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

func RemoveCopy[T comparable](r1, r2 []T, first, last, d_first int, value T) int {
	for ; first != last; first++ {
		if r1[first] != value {
			r2[d_first] = r1[first]
			d_first++
		}
	}
	return d_first
}

func RemoveCopyIf[T any](r1, r2 []T, first, last, d_first int, p func(T) bool) int {
	for ; first != last; first++ {
		if !p(r1[first]) {
			r2[d_first] = r1[first]
			d_first++
		}
	}
	return d_first
}

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

func Reverse[T any](r []T, first, last int) {
	for last--; first < last; {
		IterSwap(&r[first], &r[last])
		first++
		last--
	}
}

func ReverseCopy[T any](r1, r2 []T, first, last, d_first int) int {
	for ; first != last; d_first++ {
		last--
		r2[d_first] = r1[last]
	}
	return d_first
}

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

func RotateCopy[T any](r1, r2 []T, first, n_first, last, d_first int) int {
	return Copy(r1, r2, first, n_first, Copy(r1, r2, n_first, last, d_first))
}

func ShiftLeft[T any](r []T, first, last, n int) int {
	if n == 0 {
		return last
	}

	if n >= last-first {
		return first
	}

	return Move(r, r, first+n, last, first)
}

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
