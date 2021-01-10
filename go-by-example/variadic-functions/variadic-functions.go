package main

import "fmt"

// This is a function that will accept any number of integers as parameters
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    // If you already have multiple args in a slice, apply them to
    // a variadic function using func(slice...) like this.
    // The ... unpacks the variable into its constituent parts
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
