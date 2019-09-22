// Package strain provides filter and inverse filters for collections
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns a new list of items which match the provided predicate
func (xs Ints) Keep(filter func(int) bool) Ints {
	result := []int{}
	for _, x := range xs {
		if filter(x) {
			result = append(result, x)
		}
	}

	if len(result) > 0 {
		return result
	} else {
		return nil
	}
}

// Discard returns a new list of items which do not match the provided predicate
func (xs Ints) Discard(filter func(int) bool) Ints {
	return xs.Keep(func(x int) bool { return !filter(x) })
}

func (xs Lists) Keep(filter func([]int) bool) Lists {
	result := [][]int{}
	for _, x := range xs {
		if filter(x) {
			result = append(result, x)
		}
	}

	if len(result) > 0 {
		return result
	} else {
		return nil
	}
}

func (xs Strings) Keep(filter func(string) bool) Strings {
	result := []string{}
	for _, x := range xs {
		if filter(x) {
			result = append(result, x)
		}
	}

	if len(result) > 0 {
		return result
	} else {
		return nil
	}
}
