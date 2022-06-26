package main

import "fmt"

func main() {
	//var a int
	//var b int
	//
	//a = 1
	//b = 2

	//var a float32
	//var b float32
	//a = 1.0
	//b = 2.0

	//a := 1.0
	//b := 2.0
	a, b := 1.0, 2.0
	mean := (a + b) / 2
	fmt.Printf("a=%v,type of %T ", a, a)
	fmt.Printf("b=%v,type of %T ", b, b)
	fmt.Printf("mean=%v,type of %T ", mean, mean)
}
