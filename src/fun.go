package fun

func map_keys[K comparable, V any](m map[K]V) []K {
	out := []K{}
	for key := range m {
		out = append(out, key)
	}
	return out
}

func map_values[K comparable, V any](m map[K]V) []V {
	out := []V{}
	for _, val := range m {
		out = append(out, val)
	}
	return out
}

func map_include[K comparable, V any](m map[K]V, key K) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}

func ternary[T any](expr bool, yay, nay T) T {
	if expr {
		return yay
	}
	return nay
}
