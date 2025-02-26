package pyutil

import (
	"fmt"
)

func Len(s any) int64 {
	return int64(len(fmt.Sprint(s)))
}

func Str(s any) string {
	return fmt.Sprint(s)
}
