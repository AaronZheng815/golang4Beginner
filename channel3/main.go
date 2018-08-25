package main

import (
	// "time"
	"fmt"
)

func calc(taskChan chan int, rstChan chan int, exitChan chan bool){
	for v := range taskChan{
		flag := true
		for i :=2; i< v; i++{
			if(v%i ==0){
				flag = false
				break
			}
		}

		if flag == true{
			rstChan <- v
		}
	}
	exitChan <- true
}

func main()  {
	intChan := make(chan int,1000)
	rstChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go func(){            //匿名函数
		for i:=0;i<10000;i++{
			intChan <- i
		}
		close(intChan)  //输入管道完毕后关闭，避免range取管道出现异常
	}()

	for i:=0; i<8; i++{
		go calc(intChan, rstChan, exitChan)
	}	

	//等待所有的goroute全部退出
	go func(){
	for i:=0;i<8;i++{
		<- exitChan
		// a := <- exitChan
		fmt.Printf("goroute %d exit\n",i)
	}
	close(rstChan)
	}()
	
	// go func(){
	for v := range rstChan {  //取管道的数据，当检测到chan关闭后，取完数据
		fmt.Println(v)		
	}
	// }()

	// time.Sleep(time.Second*10)
}