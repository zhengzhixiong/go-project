package main

import (
	"log"
	"net"
	"time"
)

// 创建服务端监听
func CreateListen(serverAddr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	tcpListen, err := net.ListenTCP("tcp", tcpAddr)
	return tcpListen, err
}
func main() {
	tcpListener, err := CreateListen(":8098")
	if err != nil {
		panic(err)
	}
	log.Printf("[服务端] 监听中：%v\n", tcpListener.Addr().String())

	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("TcpListener Accept TCP Error:%v\n", err)
			return
		}
		go KeepAlive(conn)
	}
}

// 发送心跳包
func KeepAlive(conn *net.TCPConn) {
	for {
		_, err := conn.Write([]byte("KeepAlive\n"))
		if err != nil {
			//客户端异常关闭 write tcp 127.0.0.1:8098->127.0.0.1:64855: write: broken pipe
			log.Printf("[KeepAlive] Error %s", err)
			return
		}
		time.Sleep(time.Second * 3)
	}
}
