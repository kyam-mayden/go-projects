package main

import "fmt"

func plus(a int, b int) int {

    // Go requires explicit returns, whereas a language like Rust will
    // automatically return the last statement.
    return a + b
}

// When you have multiple consecutive params with the same type you can just hint the
// last one, like so:
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    // Call a function with name(args)
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
