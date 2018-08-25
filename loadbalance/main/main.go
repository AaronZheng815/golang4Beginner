package main

import (
	"time"
	"os"
	"math/rand"
	"fmt"
	"basic/loadbalance/loadbalance"
)

func main()  {
	//insts := make([]*loadbalance.Instance)
	var insts [] *loadbalance.Instance
	for i:=0; i<5; i++{
		host := fmt.Sprintf("192.168.%d,%d",rand.Intn(255), rand.Intn(255))
		inst := loadbalance.NewInstance(host,8080)
		insts = append(insts,inst)
	}

	fmt.Println(insts)

/*	var lb loadbalance.LoadBalance
	var conf = "random"
	if len(os.Args) > 1{
		conf = os.Args[1]
	}

	if conf == "random"{
		lb = &loadbalance.RandomBalance{}
		fmt.Println("use random balancer")
	} else if conf == "roundrobin"{
		lb = &loadbalance.RoundRobinBalance{}
		fmt.Println("use roundrobin balancer")
	}

	for {
		inst,err := lb.DoBalance(insts)
		if err != nil {
			fmt.Println("do balance error: ", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}*/

	var conf = "random"
	var key = "hashkey"
	if len(os.Args) > 1{
		conf = os.Args[1]
		if len(os.Args) > 2 {
			key = os.Args[2]
		}
	}
	conf = "hash"
	key = "aaron"
	for{
		inst,err := loadbalance.DoLB(conf, insts, key)
		if err != nil {
			fmt.Println("do balance error: ", err)
			continue
		}

		fmt.Printf("Host %s, port %d\n",inst.GetHost(), inst.GetPort())
		time.Sleep(time.Second)
	}
}