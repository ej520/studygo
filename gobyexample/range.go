//range关键字在for循环中用于遍历数组(array/string)，切片(slice)，通道(chan)或映射(map)的项目.
//array&&slice: return(index,value)
//string: return(index,value)
//map: return(key ,value)
//chan: return(element, none)

package main

import (
	"fmt"
)

func main(){
	//1. create a slice 
	var s = []int{0,2,4,6,8}

	////print slice using keys
	for i := range s {
		fmt.Printf("slice s[%d] is %d\n",i,s[i])
	}

	//print slice using key-value
	for i,v := range s {
		fmt.Printf("slice index is %d, slice s[%d] is %d\n",i,i,v)
	}

	//2. create a map
	var m = map[string]string{"wang":"bo","wyj":"ej520","boy":"baby"}

	//print map using keys
	for k :=range m {
		fmt.Println("key is",k,"value is",m[k])
	}


	//print map using key-value
	for k,v := range m {
		fmt.Println("key is",k,"value is",v);
	}
}