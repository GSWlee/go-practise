package main

import (
	"fmt"
	"strings"
)

func comma1(s string) string {
	n:=len(s)
	if n<=3{
		return s
	}else {
		return comma1(s[:n-3])+","+s[n-3:]
	}
}

func comma2(s string) string {
	n:=len(s)
	if n<=3{
		return s
	}else {
		return s[:3]+","+comma2(s[3:])
	}
}

func comma(s string) string {
	n:=len(s)
	if n<3{
		return s;
	}else {
		point:=strings.Index(s,".")
		return comma1(s[:point])+"."+comma2(s[point+1:])
	}
}

func main()  {
	ans:="123456789.987654321"
	a:=comma(ans)
	fmt.Println(a)
}
