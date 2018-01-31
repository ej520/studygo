//延迟函数调用被压入一个栈中。 当函数返回时，会按照后进先出的顺序调用被延迟的函数调用。

package main

import "fmt"

func main(){
	fmt.Println("continue")

	for i:=0;i<10;i++{
		defer fmt.Println(i)   // 0,1,2..9 入栈；最后打印：9..2,1,0
		fmt.Println("this number is:",i)
	}

	fmt.Println("done")
}


