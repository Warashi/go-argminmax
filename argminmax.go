package argminmax

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i is less than the element with index j.
	Less(i, j int) bool
}

func ArgMin(data Interface) int {
	argmin := 0
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, argmin) {
			argmin = i
		}
	}
	return argmin
}

func ArgMax(data Interface) int {
	argmax := 0
	for i := 1; i < data.Len(); i++ {
		if data.Less(argmax, i) {
			argmax = i
		}
	}
	return argmax
}

func ArgMinf(length int, less func(i, j int) bool) int {
	argmin := 0
	for i := 1; i < length; i++ {
		if less(i, argmin) {
			argmin = i
		}
	}
	return argmin
}

func ArgMaxf(length int, less func(i, j int) bool) int {
	argmax := 0
	for i := 1; i < length; i++ {
		if less(argmax, i) {
			argmax = i
		}
	}
	return argmax
}
