package main

import (
	"errors"
	"time"
	"fmt"
)

func initConfig() (err error){
	return errors.New("init config failed")
}

func test()  {
	/*defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()*/

    b := 0
	a := 100/b
	fmt.Println(a)
	
	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}
func main(){
	var str1 = "hello world\n\n"
	var str2 = `hello world\n
	this is a test string\n
	this is a test string`

	fmt.Println("str1 = ", str1)
	fmt.Println("str2 = ", str2)

	intPtr := new(int32)
	*intPtr = 123
	fmt.Println(*intPtr)
	fmt.Println(intPtr)

	var arr []int
	arr = append(arr, 10,20,234)
	fmt.Println(arr)

	arr = append(arr, arr...)
	fmt.Println(arr)

	for{
		test()
		time.Sleep(time.Second)
	}
	
}