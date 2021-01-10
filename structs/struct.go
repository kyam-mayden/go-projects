package main

import "fmt"

// Structs are typed collections of fields. Like OOPs objects but without methods attached
// This is a person struct with 2 fields: name and age
type person struct {
    name string
    age  int
}

// newPerson will return a Person struct with the supplied name
func newPerson(name string) *person {

    p := person{name: name}
    p.age = 42
    return &p
}

func main() {
    // This creates a new person
    fmt.Println(person{"Bob", 20})

    // This creates a new person with named parameters
    fmt.Println(person{name: "Alice", age: 30})

    // Any omitted fields will be zero-valued (which varies depending on type)
    fmt.Println(person{age: "Fred"})

    // & will return the pointer to a struct
    fmt.Println(&person{name: "Ann", age: 40})

    // Good practice to encapsulate new struct creation in a constructor function
    fmt.Println(newPerson("Jon"))

    // Access struct fields with dot syntax
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // You can access the values from a struct even if you only have the pointer reference
    sp := &s
    fmt.Println(sp.age)

    // Structs are mutable
    sp.age = 51
    fmt.Println(sp.age)
}
