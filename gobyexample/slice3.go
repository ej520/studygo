//slice apppend() copy()
package main

import (
	"fmt"
)

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func main(){
	var s []int
	printSlice(s)
	//添加1个元素
	s = append(s,0)
	printSlice(s)
	//添加多个元素
	s = append(s,1,2,3,4)
	printSlice(s)

	//create new slice 
	s1:= make([]int,len(s),cap(s)*2)
	printSlice(s1)
	//copy s to s1
	copy(s1,s)
	printSlice(s1)
}