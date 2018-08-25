package main

import (
	"time"
	"fmt"
)

func main()  {
	ch := make(chan int,100)
	ch2 := make(chan int,100)

	go func(){
		for i:=0;i<100;i++{
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i
		}
	}()
	
	for {
		// data := <- ch
		// fmt.Println(data)
		//当读完数据后，会死锁

		//通过select可以解决这个问题
		select {
		case data := <- ch:
			fmt.Println(data)
		case data := <- ch2:
			fmt.Println(data)
		default:
			fmt.Println("No data in the channel!")
			time.Sleep(time.Second)
		}
	}
}