package main

import (
	"fmt"
)

func Test(a interface{}) {
	tmp := a.(int)
	tmp += 3
	fmt.Println(tmp)
}

func Test1(a interface{}){
	tmp, ok := a.(int)
	if ok == false{
		fmt.Println("Convert to int failed")
		return
	}
	tmp += 3
	fmt.Println(tmp)
}


func main(){
	// var a interface{}

	// var b int = 1
	// a = b
	// c := a.(int)  //接口转换成Int类型

	// fmt.Printf("%d %T\n",a,a)
	// fmt.Printf("%d %T",c,c)

	var b int
	Test(b)

	var b1 string
	Test1(b1)

}