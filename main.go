package main

import (
	"fmt"

	"github.com/kallol7/go-py-util/pyutil"
)

func main() {
	strLen := pyutil.Len("résumé")
	fmt.Println("Length of text:", strLen) // Length of text: 6

	floatToStr := pyutil.Str(3.141592653589793)
	fmt.Printf("Number: %v, Type: %[1]T", floatToStr) // Number: 3.141592653589793, Type: string
}
