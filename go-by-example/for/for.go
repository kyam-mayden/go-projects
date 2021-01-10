package main

import "fmt"

func main() {
	// for is Go’s only looping construct.
	i := 1
	// Basic, with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// classic initial / condition/ after
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}