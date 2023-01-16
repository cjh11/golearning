package main

import (
	"fmt"
	"math"
)

func main() {
	//局部声明变量方式1
	var a int
	fmt.Println("a=", a)
	//声明变量方式2
	var b int = 100
	fmt.Println("b=", b)
	//声明变量方式3
	var c = 200
	fmt.Println("c=", c)
	fmt.Printf("c=%v,type of c=%T\n", c, c)
	//声明变量方式4(常用)
	e := 300
	fmt.Printf("e=%v,type of e=%T\n", e, e)
	f := float64(e)
	fmt.Printf("f=%f,type of f=%T\n", f, f)

	//常量
	const eh = 2000
	const i = eh / 4
	fmt.Println(eh, math.Sin(i))

	var ll, yy = 123, "sad"
	fmt.Println(ll, yy)
}
