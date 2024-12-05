package internal

import (
	"fmt"
	"strconv"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func ToInt(d string) int {
	v, err := strconv.Atoi(d)
	if err != nil {
		panic(fmt.Sprintf("Unexpected character (%s). Expected digit.", d))
	}
	return v
}
