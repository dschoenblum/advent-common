package common

func AnyKey[K comparable, V any](m map[K]V) K {
	for key := range m {
		return key
	}
	panic("no keys found")
}
