package main

import "fmt"

var p *int
func main(){

	i,j := 10,3210
	p = &i

	fmt.Println(p)
	fmt.Println(*p)

	*p = 20
	fmt.Println(i)

	p = &j
	fmt.Println(p)
	fmt.Println(*p)
	*p = *p/10
	fmt.Println(j)
}