//go循环只有for循环
package main

import (
	"fmt"
)

func main(){
/*	sum:=0
	for i:=1;i<=100 ; i++ {
		sum=sum+i
	}

	fmt.Println(sum)
	*/

	/* 初始化语句和后置语句是可选的
	sum ,i:= 0 ,1
	for ;i<=100;{
		sum+=i
		i++
	}
	fmt.Println(sum)
	*/

	//当做while使用
	sum := 1
	for sum<=5050 {
		sum+=sum
	}
	fmt.Println(sum)

	//死循环
	for {//无退出条件，变成死循环
		fmt.Println("dead for")
	}
}



