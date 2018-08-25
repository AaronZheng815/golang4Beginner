package main

import (
	"fmt"
	"reflect"
)

type Student struct{
	Name string
	Age int
}

func test(b interface{}){
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	fmt.Println(v)

	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu,ok:=iv.(Student)
	if ok {
		fmt.Printf("%v %T\n",stu,ok)
	}
}
func testInt(b interface{}){
	val :=reflect.ValueOf(b)
	// c := val.Int()
	c :=val.String()
	fmt.Printf("get value of interface %s\n", c)
}
func setInt(i interface{}){
	val := reflect.ValueOf(i)
	val.Elem().SetInt(300)
}

func main()  {
	var a int = 200
	test(a)	
	testInt(a)

	stu :=Student{"aaron", 18}
	test(stu)

	b:=100
	setInt(&b)
	fmt.Println(b)

}