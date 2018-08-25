package main

import (
	"fmt"
)

type Student struct{
	Name string
	Age int
}

func main()  {
	//传递整型
	var ch1 chan int
	ch1 = make(chan int, 10)
	ch1 <- 10

	//传递Map
	var mapch chan map[string]string
	mapch = make(chan map[string]string, 10)
	m := make(map[string]string,16)
	m["stu01"] = "aaron"
	m["stu02"] = "nina"
	mapch <- m

	//传递Struct
	var structch chan Student
	structch = make(chan Student, 10)
	stu := Student{
		Name :"aaron",
		Age :18,
	}
	structch <- stu
	var stu2 Student
	stu2 = <- structch
	fmt.Println(stu2)

	//传递Interface
	var stuch chan interface{}
	stuch = make(chan interface{},10)

	stuch <- &stu  //任何类型都可以赋值给空接口，这里用指针来测试
	var stu3if interface{}
	stu3if = <- stuch

	var stu3p *Student
	stu3p,ok :=stu3if.(*Student)
	if !ok {
		fmt.Println("Cannot convert.")
		return
	}
	fmt.Println(stu3p,*stu3p)
}