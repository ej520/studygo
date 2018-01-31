package main

import (
	"time"
	"fmt"
)

func main(){
	fmt.Println(time.Now())
	time1 := time.NewTimer(time.Second*2)
	<-time1.C
	fmt.Println(time.Now(),time1)


	time2 := time.NewTimer(time.Second*5)
	go func(){
		<-time2.C
		fmt.Println("time2 expired")
	}()
	
	
}