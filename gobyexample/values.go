package main

import (
	"fmt"
	"math/rand"
	"math"
)

func split(sum int) (x,y int){
	x=sum - 10
	y=sum - x
	return y ,x 
}

func main(){
	fmt.Println("go" + "lang")

	fmt.Println("1+1=",1+1)
	fmt.Println("7.0/3.0=",7/3)

	//运行环境时确定的，rand.Intn每次会返回相同的数字
	fmt.Println("my rand number is :" , rand.Intn(100))

	fmt.Println("4的平方根：",math.Sqrt(4))

	fmt.Println(split(25))
}
