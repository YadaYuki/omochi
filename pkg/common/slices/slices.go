package slices

import "math"

// ref: https://pkg.go.dev/golang.org/x/exp/slices

func Contains[T comparable](slice []T, tgt T) bool {
	for _, item := range slice {
		if item == tgt {
			return true
		}
	}
	return false
}

func SplitSlice[T any](slice []T, size int) [][]T {
	var splitedSlices [][]T
	for i := 0; i < len(slice); i += size {
		tail := math.Min(float64(len(slice)), float64(i+size))
		splitedSlices = append(splitedSlices, slice[i:int(tail)])
	}
	return splitedSlices
}
