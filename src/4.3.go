package main

import "fmt"

func reverse(ptr *[10]int)  *[10]int{
	for i,j:=0,9;i<j;i,j=i+1,j-1{
		ptr[i],ptr[j]=ptr[j],ptr[i]
	}
	return ptr
}
func main()  {
	num:=[10]int{0,1,2,3,4,5,6,7,8,9}
	var nums *[10]int =&num
	nums=reverse(nums)
	fmt.Println(nums)
}