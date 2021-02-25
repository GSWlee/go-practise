package main

import (
	"fmt"
)

const boilingF1,boilingF2 =212.0,110.0

func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	d := [2]int{1, 2}
	fmt.Println(a == d)
}
