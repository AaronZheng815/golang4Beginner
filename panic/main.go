package main

import (
	"fmt"
	"time"
)

func test(){
	//通过revocer捕获了错误，不会影响其它协程
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("panic: ",err)
		}
	}()
	
	var m map[string]int
	m["student"] = 100

}

func calc(){
	for {
		fmt.Println("I am Working...")
		time.Sleep(time.Second)
	}
}

func main()  {
	for i:=0;i<100;i++{
		go test()
	}
	go calc()

	time.Sleep(time.Second*10)	
}