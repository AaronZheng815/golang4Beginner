package tcpSer

import (
	"fmt"
	"net"
)

type TcpServer struct{
	Ip string
	Port int
}

func NewTcpServer(ip string, port int) (tcpServer *TcpServer, err error)  {
	tcpServer = &TcpServer{
		Ip : ip,
		Port : port,
	}
	err = nil
	return
}

func (p *TcpServer) Start() (err error){
	address := fmt.Sprintf("%s:%d",p.Ip,p.Port)

	listen, err := net.Listen("tcp",address)
	if err != nil{
		fmt.Println("tcp listen failed, error:",err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept failed, error:",err)
			continue
		}
		go msgHandler(conn)
	}
}

func msgHandler(conn net.Conn){
	//通过revocer捕获了错误，不会影响其它协程
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("panic: ",err)
		}
	}()

	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		rc, rerr:= conn.Read(buf)
		if rerr != nil {
			fmt.Println("Read error:", rc, rerr)
			conn.Close()
			return
		}
		msg := string(buf)
		fmt.Println("Read from Client:", msg)

		_, werr := conn.Write(buf)
		if werr != nil {
			fmt.Println("write error:", werr)
			return 
		}
		fmt.Println("Ser send: ", msg)
	}
}