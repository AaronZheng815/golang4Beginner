package main

import (
	"fmt"
	"reflect"
)

type Student struct{
	Name string
	Age int
	Score float32
}

func (s Student) Set(name string, age int, score float32){
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print(){
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end----")
}

func TestStruct(stu interface{})  {
	val := reflect.ValueOf(stu)

	kd := val.Kind()
	if kd != reflect.Struct{
		fmt.Println("should be struct type")
		return 
	}

	fnum :=val.NumField()
	fmt.Printf("struct has %d fileds\n", fnum)
	for i:=0; i< fnum; i++{
		fmt.Printf("%d %v\n", i, val.Field(i))
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methords.", numOfMethod)

	// var params []reflect.Value
	// val.Method(0).Call(params)
	val.Method(0).Call(nil)

	fmt.Printf("%v \n", val.Method(1))
	//val.Method(1).Call(params)
}

func main()  {
	a := Student{"Aaron", 18,95.5}

	TestStruct(&a)
	
}