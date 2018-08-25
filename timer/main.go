package main

import (
	"time"
	"fmt"
)

func main()  {
	t1 := time.NewTicker(time.Second)
	t2 := time.After(time.Second*2)
	var count int
	for {
		select{
		case v := <- t1.C:
			fmt.Println(v.Local())
			count ++
			if count >10 {
				t1.Stop()
				break
			}
		case v:= <- t2:  //不建议使用After定时器，会内存泄漏
			fmt.Println("after timer: ", v.Local())
		default:
			//fmt.Println("No event.")
		}
	}
}