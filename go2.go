package main

import (
	. "constraints"
   	"fmt"
)

func Max[T Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func main() {
	m := Max(1, 2)
	fmt.Println(m)
}
