//switch case default fallthrough 

package main

import (
	"time"
	"fmt"
)

func main(){
	fmt.Println("GO run switch:")
	
	switch i:=200 ; {
	case i>10: fmt.Println("i>10")
	 //有fallthrough匹配结果后，强制执行下一个分支
	fallthrough;                  
	case i>500: 
		fmt.Println("i>500",i)   //强制执行该分支，并不做条件判断
	case i>100: fmt.Println("i>100")
	default:
		fmt.Println(i)
	} 

	//无条件的switch，替代if-else链，更清晰
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

}
