package test

import (
	"octopus/tcp"
	"octopus/udp"
	"testing"
	"time"
)

func TestUdp(t *testing.T) {
	udpStart := udp.UdpStart{Port: 40000}
	udpStart.Start()
	udpMsg := udp.UdpMsg{Ip: "127.0.0.1", Port: 30000, Msg: "mc_cao"}
	udp.SendMsg(&udpMsg)
	time.Sleep(1 * time.Second)
}

func TestTcp(t *testing.T) {
	tcpStart := tcp.TcpStart{TcpClinetPort: 40000}
	tcpStart.Start()
	tcpMsg := tcp.TcpMsg{Ip: "127.0.0.1", Port: 20000, Msg: "mc_cao"}
	tcp.SendMsg(&tcpMsg)
	time.Sleep(1 * time.Second)
	//conn, err := net.Dial("tcp", "127.0.0.1:20000")
	//if err != nil {
	//	fmt.Println("err : ", err)
	//	return
	//}
	//defer conn.Close() // 关闭TCP连接
	//for {
	//	input :="mc_cao" // 读取用户输入
	//	bytes:=[]byte(input)
	//	fmt.Println(bytes)
	//	_, err := conn.Write(bytes) // 发送数据
	//	if err != nil {
	//		return
	//	}
	//	buf := [512]byte{}
	//	n, err := conn.Read(buf[:])
	//	if err != nil {
	//		fmt.Println("recv failed, err:", err)
	//		return
	//	}
	//	fmt.Println(string(buf[:n]))
	//}
}
