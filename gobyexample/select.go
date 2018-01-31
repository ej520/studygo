//select gorouting channel

package main

import (
	"time"
	"log"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(time.Second*2)
		c1 <- "one"
		time.Sleep(time.Second*4)
		c1 <- "oneone"
	}()

	go func(){
		time.Sleep(time.Second*2)
		c2 <- "two"
		time.Sleep(time.Second*8)
		c2 <- "two222"
	}()

	log.Println(time.Second)
	for i:=0; i<4;i++ {
		select{
		case msg1 := <-c1:
			log.Println("recieved:",msg1)
			
		case msg2:=<-c2:
			log.Println("received:"+msg2)
		}
	}
}