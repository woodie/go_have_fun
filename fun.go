package fun

func ternary[T any](expr bool, yay, nay T) T {
	if expr {
		return yay
	} else {
		return nay
	}
}
