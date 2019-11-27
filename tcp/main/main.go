package main

import (
	"github.com/AaronZheng815/golang4Beginner/tcp/tcpClt"
	"github.com/AaronZheng815/golang4Beginner/tcp/tcpSer"
	"fmt"
	"flag"
)

func main()  {
	//lens :len(os.Args)
	var role string
	flag.StringVar(&role, "role","clt","Please input role{ser,clt}")

	var ip string
	flag.StringVar(&ip, "ip","127.0.0.1","Please input Server Ipv4 address")

	var port int
	flag.IntVar(&port,"port",50000,"Please input port")

	flag.Parse()
	fmt.Printf("role: %s, ip address: %s, port: %d\n", role, ip, port)	
	
	if role == "ser"{
		tcpServer, err := tcpSer.NewTcpServer(ip, port)
		if err != nil {
			fmt.Printf("Failed to create tcp server, error: %v\n", err)
			return
		}
		tcpServer.Start()
	} else{
		tcpClient, err := tcpClt.NewTcpClient(ip,port)
		if err != nil {
			fmt.Printf("Failed to create tcp client, error: %v\n", err)
			return
		}
		tcpClient.StartEcho()
	}
	return
}
