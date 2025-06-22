package pyutil

import (
	"fmt"
)

func Len(s any) int64 {
	if r, ok := s.(string); ok {
		return int64(len([]rune(r)))
	}

	return Len(fmt.Sprint(s))
}

func Str(s any) string {
	return fmt.Sprint(s)
}
