package main

import (
	"fmt"

	"github.com/kallol7/go-py-util/pyutil"
)

func main() {
	fmt.Println("Length of text:", pyutil.Len("bigtext"))
}
