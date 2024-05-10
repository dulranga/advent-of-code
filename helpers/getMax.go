package helpers

import "cmp"

func GetMax[T cmp.Ordered](slice []T) T {
	var max T
	for _, item := range slice {
		if item > max {
			max = item
		}
	}

	return max
}
