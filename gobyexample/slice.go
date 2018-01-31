//slice array make 
//len(slice1) cap(slice1)

package main

import (
	"fmt"
)



func main(){
	var num1 []int
	fmt.Println(num1)
	num1 = make([]int ,5,5)  //初始化方法1(类型、长度5、容量5)，默认5个元素全为0
	num1[1] = 5
	fmt.Println(num1)


	var num2 = []int{1,2,3,4,5} //初始化方法2
	fmt.Println(num2)

	printSlice(num1)

	var num3 []int  //缺省情况，slice is nil ，len =0, cap=0
	printSlice(num3)
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
