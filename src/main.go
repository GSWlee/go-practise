package main

import "fmt"

type X struct {
	a int
	b int
}

type Y struct {
	X
	a int
}

func main() {
	y :=Y{X{1,2},3}
	fmt.Println(y)
}

