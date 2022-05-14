package udp

import "fmt"

//Udp停止方法
type UdpStop struct {
}

func (d *UdpStart) Stop() {
	getInstance().Stop()
	fmt.Println("UdpStop stop")
}
