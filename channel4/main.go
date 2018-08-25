package main

import (
	"fmt"
)
//传只写的管道
func sendData(sndChan chan<- string, exitChan chan bool)  {
	for i:=0;i < 1000;i ++{
		sndChan <- fmt.Sprintf("hello world %d", i)
	}
	close(sndChan) //需要关闭，否则recv那边会阻塞
	exitChan <- true
}

//传只读的管道
func recvData(sndChan <-chan string, exitChan chan bool){
	for {
		v,ok := <- sndChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
	exitChan <- true
}

func main()  {
	dataChan := make(chan string, 100)
	exitChan := make(chan bool, 2)

	go sendData(dataChan,exitChan)
	go recvData(dataChan,exitChan)

	for i:=0;i<2;i++{
		<- exitChan
		fmt.Println("goroute exit")
	}

}