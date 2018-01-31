//if 就像for 一样， 不要求用(), 必须有{}
package main

import (
	"math"
	"fmt"
)

func sqrt(x float64) string{
	if x<0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x,y,lim float64) float64 {
	if v:=math.Pow(x,y); v<lim{
		return v
	}else{
		fmt.Println(v,lim)
	}
	return lim
}

func main(){
	fmt.Println(sqrt(-10),sqrt(10))
	
	fmt.Println(
		pow(3,4,100),
		pow(3,4,10),
		pow(3,4,100),
		pow(3,4,100),
	)
}