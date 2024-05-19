package main

import "fmt"

const (
	typedConstant   = int16(100)
	untypedConstant = 100
)

func main() {
	i := int(1)
	fmt.Println("unTyped:", i*untypedConstant)

	// should be an error types mismatch
	fmt.Println("Typed:", i*typedConstant)
}
