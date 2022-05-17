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
	tcpStart := tcp.TcpStart{ServerPort: 40000}
	tcpStart.Start()
	tcpMsg := tcp.TcpMsg{Ip: "127.0.0.1", Port: 20000, Msg: "mc_cao"}
	tcp.SendMsg(&tcpMsg)
	time.Sleep(1 * time.Second)
}
