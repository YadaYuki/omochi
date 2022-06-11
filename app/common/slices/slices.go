package slices

// ref: https://pkg.go.dev/golang.org/x/exp/slices

func Contains[T comparable](slice []T, tgt T) bool {
	for _, item := range slice {
		if item == tgt {
			return true
		}
	}
	return false
}
