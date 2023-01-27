package utils

import "context"

func GetIntFromContext(ctx context.Context, key string, defaultValue int) int {
	vI := ctx.Value(key)
	v, ok := vI.(int)
	if !ok {
		return defaultValue
	}

	return v
}

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func DeleteAtIndex[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
