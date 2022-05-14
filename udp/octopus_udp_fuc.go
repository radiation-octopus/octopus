package udp

import (
	"sync"
)

var octopusUdp *OctopusUdp

var once sync.Once

//单例模式
func getInstance() *OctopusUdp {
	once.Do(func() {
		octopusUdp = new(OctopusUdp)
	})
	return octopusUdp
}

func SendMsg(udpMsg *UdpMsg) {
	getInstance().sendMsg(udpMsg)
}

//
//func UdpAcceptCallBinding(method string){
//	UdpAcceptCallBindingMethod=method
//}
