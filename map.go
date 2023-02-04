package common

func AnyKey[K comparable, V any](m map[K]V) K {
	for k := range m {
		return k
	}
	panic("no keys found")
}

func AddMap[K comparable, V any](dst map[K]V, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}

func CloneMap[K comparable, V any](src map[K]V) (dst map[K]V) {
	dst = make(map[K]V, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return
}
