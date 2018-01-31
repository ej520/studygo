/*
 variables:
 bool
 string
 int  int8 int16 int32 int64
 uint uint8 uint16 uint32 uint64 uinitptr
 
 byte => uinit8  alias
 rune //int32 alias  // unioncode 

 float32 float64
 complex64 complex128

 // int uint uintptr 32bit os : 32bit    64bit os 64bit 整数首选int  特别的理由采用
*/

package main
import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	tobe bool = false
	maxint uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5+12i)
	
)


func main(){
	const f = "%T(%v)\n"

	fmt.Printf(f,tobe,tobe)
	fmt.Printf(f,maxint,maxint)
	fmt.Printf(f,z,z)

	var a bool
	var b int
	var c string
	fmt.Println("a,b,c:",a,b,c)
	fmt.Printf(f,a,a)
	fmt.Printf(f,b,b)
	fmt.Printf(f,c,c)

	//类型转换
	var i int = 10
	var fl float64 = float64(i)
	fmt.Printf(f,i,i)
	fmt.Printf(f,fl,fl)

	var x,y int = 3, 6
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	fmt.Printf(f,f1,f1)

	const big = 1<<100
	fmt.Printf(big)
}


