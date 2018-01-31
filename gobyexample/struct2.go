package main

import "fmt"

type Vertex struct{
	x, y int
}

var (
	a = Vertex{1,2}
	b = Vertex{x:1}
	c = Vertex{}
	p = &Vertex{1,2}
)

func main(){

	fmt.Println(a,b,c,p)
}