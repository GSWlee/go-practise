package main

import (
	"./length"
	"fmt"
	"os"
	"strconv"
)

const boilingF1,boilingF2 =212.0,110.0

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := length.Meter(t)
		k := length.Kilometer(t)
		fmt.Printf("%g = %g, %g = %g\n",
			k, length.KToM(k), m, length.MToK(m))
	}
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}