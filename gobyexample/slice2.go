//slice 切片 子切片


package main

import (
	"fmt"
)

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main(){
	//create a slice
	var s = []int{0,1,2,3,4,5}
	printSlice(s)

	//print the original slice
	fmt.Println("slice ==",s)

	//s[1,4] 打印index 1 到3，不包含4
	fmt.Println("slice[1,4] ==",s[1:4])
	//s[:4] 打印index 0 到3, 不包含4
	fmt.Println("slice[:4] ==",s[:4])
	//s[1:] 打印index 从1到max
	fmt.Println("slice[1:] ==", s[1:])

	var s1 = make([]int,0,5)
	printSlice(s1)

	s2 := s[:4]
	printSlice(s2)
	s3 := s[1:4]
	printSlice(s3)
}