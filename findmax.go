package mcommon

import (
	"errors"
)

// Max finds the maximum given the provided 'more' function.
func Max(slice []interface{}, more func(i, j int) bool) (interface{}, error) {
	if len(slice) == 0 { return nil, errors.New("Cannot find maximum for empty slice") }
	if len(slice) == 1 { return slice[0], nil }
	j := 0
	for i := 1; i < len(slice); i++ {
		if more(i, j) {
			j = i
		}
	}
	return slice[j], nil
}
