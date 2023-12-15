package utility

type Pair[T1, T2 any] struct {
	first  T1
	second T2
}

func MakePair[T1, T2 any](t T1, u T2) Pair[T1, T2] {
	return Pair[T1, T2]{t, u}
}
