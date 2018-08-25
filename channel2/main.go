package main

import (
	"fmt"
	"time"
)


func write(ch chan int){ //chan是引用类型，不需要传指针
	for i :=0;i <200;i ++{
		ch <- i
		fmt.Println("write into channel: ", i)
	}
}

func read(ch chan int){
	for {
		var rch int
		rch = <- ch
		time.Sleep(time.Millisecond*50)
		fmt.Println("read from channel: ", rch)
	}
}

func main(){
	//var intch chan int
	intch := make(chan int, 50)
	
	go write(intch)
	go read(intch)

	time.Sleep(10*time.Second)	
}