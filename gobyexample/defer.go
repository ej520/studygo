//defer 语句会延迟函数的执行，直到上层函数返回。
package main

import (

	"log"
)

func main(){
	defer log.Println("defer: 3How are you!")

	defer log.Println("defer: 2Hello!!")

	log.Println("1print")
}