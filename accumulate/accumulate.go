// package accumulate provides an methods for mapping strings
package accumulate

// AccumulateInPlace applies the transformation t to all strings of the input slice, mutating the input.
func AccumulateInPlace(xs []string, t func(string) string) []string {
	for i, x := range xs {
		xs[i] = t(x)
	}

	return xs
}

// AccumulateInPlace applies the transformation t to all strings of the input slice, returning a new slice with the resulting values
func Accumulate(xs []string, t func(string) string) []string {
	result := make([]string, len(xs))
	copy(result, xs)
	return AccumulateInPlace(result, t)
}
