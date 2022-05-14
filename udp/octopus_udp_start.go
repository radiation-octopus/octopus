package udp

import "fmt"

//Udp启动方法
type UdpStart struct {
	Port          int    `autoInjectCfg:"octopus.udp.port"`
	BindingMethod string `autoInjectCfg:"octopus.udp.binding.method"`
	BindingStruct string `autoInjectCfg:"octopus.udp.binding.struct"`
}

func (u *UdpStart) Start() {
	UdpAcceptCallBindingMethod = u.BindingMethod
	UdpAcceptCallBindingStruct = "*" + u.BindingStruct
	getInstance().Start(u.Port)
	fmt.Println("UdpStart start")
}
