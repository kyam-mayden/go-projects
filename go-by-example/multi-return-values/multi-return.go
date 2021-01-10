package main

import "fmt"
// Go can return multiple values from the same function. This is often used to return
// result and failure values

// Here, we signify that the return of the function is two integers with (int, int)
func vals() (int, int) {
    return 3, 7
}

func main() {
    // Here we use the 2 different return values from the call with multiple assignment.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}
