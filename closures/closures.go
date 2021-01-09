package main

import "fmt"

// This function will return the anonymous function, but will keep the state of i
// between calls.
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    // nextInt is the function returned from intSeq
    nextInt := intSeq()

    fmt.Println(nextInt()) // 1
    fmt.Println(nextInt()) // 2
    fmt.Println(nextInt()) // 3

    // To make sure that state doesn't carry between instances, make a new one
    newInts := intSeq()
    fmt.Println(newInts()) // 1
}
