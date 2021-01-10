package main

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean, and numeric values.
const s string = "constant"

func main() {
	fmt.Println(s)
	// A const statement can appear anywhere a var statement can.
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
