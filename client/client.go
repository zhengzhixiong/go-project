package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	conn, err := CreateConn(":8098")
	if err != nil {
		panic(err)
	}
	log.Printf("[客户端连接成功]：%v", conn.RemoteAddr().String())
	for {
		s, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Get Data Error:%v\n", err)
			continue
		} else {
			log.Printf("客户端接收到服务端数据--->%v", s)
		}

	}

}

// CreateConn 创建连接
func CreateConn(serverAddr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	return conn, err
}
