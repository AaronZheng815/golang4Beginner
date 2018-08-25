package main

import (
	"strings"
	"fmt"
)

//闭包函数
func adder() func(int) int {
	var x int
	return func(delta int) int{
		x += delta
		return x
	}
}

//闭包函数，变量为入参
func addSuffix(suffix string) func(string) string {
	return func(name string) string{
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return name
	}
}
func main()  {
	var f = adder()

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(100))

	addbmp := addSuffix(".bmp")
	addjpg := addSuffix(".jpg")
	
	fmt.Println(addbmp("test"))
	fmt.Println(addjpg("test"))

}