package main

import "fmt"

func main() {
    // Range will iterate over a set of data - and can accept different datatypes
    // Here we add up all the numbers in a slice. The first parameter is the index/key, second
    // is value. Here we don't care about the index, so it is set to _
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // Here we iterate over the same slice, but we use the index in our work.
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // Here we iterate over a map in the same way
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // Here we iterate over just the keys of a map
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // We can iterate over string literals. The Index is the position of the item
    // in the string, the value is the unicode code point of the item in decimal.
    // For example, this will output:
    //     0 103
    //     1 111
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}