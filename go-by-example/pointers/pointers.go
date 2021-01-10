package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

// Passing a value in with `*` passes it in by reference
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // The &i syntax gives the memory address of i, i.e. a pointer to i.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)
    // Pointers can be printed too
    fmt.Println("pointer:", &i)
}
