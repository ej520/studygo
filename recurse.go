//递归（recurse） 是以自相似的方式重复项的过程。在函数内调用同一个函数称为递归调用
/*
func recursion() {
   recursion() // function calls itself
}
*/

//在使用递归时，程序员需要注意在函数中定义或设置一个退出条件，否则它会进入无限循环。
//递归函数非常有用，可用于解决许多数学问题，如计算数字的阶乘，生成斐波那契数列等。
package main

import (
	"fmt"
)

//阶乘
func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

//斐波那契数列
func fibonaci(i int) int {
	if i == 0 {
		return 1
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

func main() {
	//阶乘
	var i int = 15
	fmt.Printf("factorial of %d is %d\n", i, factorial(i))

	//斐波那契数列
	for i = 0; i <= 20; i++ {
		//fmt.Printf("number %d fibonaci number is %d\n", i, fibonaci(i))
		fmt.Printf("%d ", fibonaci(i))
	}
}
