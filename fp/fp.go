package fp

// Map applies the provided function fn to each item in items
// and returns a new slice with the results.
//
// Ordering is maintained.
func Map[T, P any](fn func(T) P, items []T) []P {
	ret := make([]P, len(items))

	for i, v := range items {
		ret[i] = fn(v)
	}

	return ret
}

// Filter returns a new slice containing only the elements in items for which fn returns true.
func Filter[T any](fn func(T) bool, items []T) []T {
	var ret []T

	for _, v := range items {
		if fn(v) {
			ret = append(ret, v)
		}
	}

	return ret
}

// Reduce executes the provided function against each item in items along with an accumulator
// variable.  It returns the value of the accumulator.
func Reduce[T, P any](fn func(P, T) P, acc P, items []T) P {
	for _, v := range items {
		acc = fn(acc, v)
	}

	return acc
}
