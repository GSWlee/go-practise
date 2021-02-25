package main

import "fmt"

func main()  {
	s:=[]string{}
	s=append(s,"apple","hello","hello","hello","world","world")
	flag:=""
	i:=0
	for _,item:=range s{
		if item==flag{
			continue
		}else {
			flag=item
			s[i]=item
			i++
		}
	}
	s=s[:i]
	fmt.Println(s)
}