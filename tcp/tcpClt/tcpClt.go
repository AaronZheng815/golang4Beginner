package tcpClt

import (
	"os"
	"bufio"
	"fmt"
	"net"
	"strings"
)

type TcpClient struct{
	serIp string
	serPort int
}

func NewTcpClient(serIp string, serPort int) (tcpClient *TcpClient, err error){
	tcpClient = &TcpClient{
		serIp : serIp,
		serPort : serPort,
	}
	err = nil
	return	
}

func (p *TcpClient) StartEcho()(err error){
	serAddr := fmt.Sprintf("%s:%d",p.serIp,p.serPort)

	conn, err := net.Dial("tcp",serAddr)
	if err != nil {
		fmt.Println("Failed dialing, error:", err)
		return
	}
	
	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		//read message from stdin
		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")
		if trimInput == "quit"{
			return
		}

		_,err = conn.Write([]byte(trimInput))
		if err != nil {
			fmt.Println("Client send msg error: ",err)
			return
		}

		//read message from server
		buf := make([]byte, 1024)
		_,err = conn.Read(buf)
		if err != nil{
			fmt.Println("Client read msg error: ", err)
			return
		}
		fmt.Println("Client Read msg from Server: ", string(buf))
	}
}