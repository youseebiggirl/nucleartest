package main

import (
	"strings"
	"fmt"
)

func main() {
	s1 := "123"
	s2 := strings.Clone(s1)
	fmt.Println(s1, s2)
}
