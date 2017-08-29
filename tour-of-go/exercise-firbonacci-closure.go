package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() []int  {
	f := []int{0, 1}
	return func() []int {
		i := len(f)
		som := f[i - 1] + f[i - 2]
		f = append(f, som)
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
