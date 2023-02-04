package common

func CloneArray[V any](src []V) (dst []V) {
	dst = make([]V, len(src))
	copy(dst, src)
	return
}

func DeduplicateArray[V comparable](a []V) []V {
	m := map[V]bool{}
	for _, e := range a {
		m[e] = true
	}

	r := make([]V, len(m))
	i := 0
	for e := range m {
		r[i] = e
		i++
	}
	return r
}
