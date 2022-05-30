package udp

import (
	"github.com/radiation-octopus/octopus/log"
)

//Udp停止方法
type UdpStop struct {
}

func (d *UdpStart) Stop() {
	Stop()
	log.Info("UdpStop stop")
}
