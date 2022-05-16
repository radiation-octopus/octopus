package tcp

import "fmt"

//Tcp启动方法
type TcpStart struct {
	ServerPort           int    `autoInjectCfg:"octopus.tcp.server.port"`
	ServerBindingPoolNum int    `autoInjectCfg:"octopus.tcp.server.binding.pool.num"`
	ServerBindingMethod  string `autoInjectCfg:"octopus.tcp.server.binding.method"`
	ServerBindingStruct  string `autoInjectCfg:"octopus.tcp.server.binding.struct"`
	ClinetPort           int    `autoInjectCfg:"octopus.tcp.clinet.port"`
	ClinetMsgNum         int    `autoInjectCfg:"octopus.tcp.clinet.msg.num"`
	ClinetBindingMethod  string `autoInjectCfg:"octopus.tcp.clinet.binding.method"`
	ClinetBindingStruct  string `autoInjectCfg:"octopus.tcp.clinet.binding.struct"`
}

func (t *TcpStart) Start() {
	TcpClinetMsgNum = t.ClinetMsgNum
	TcpServerPort = t.ServerPort
	TcpServerBindingPoolNum = t.ServerBindingPoolNum
	TcpServerTcpAcceptCallBindingMethod = t.ServerBindingMethod
	TcpServerTcpAcceptCallBindingStruct = "*" + t.ServerBindingStruct
	TcpClinetPort = t.ClinetPort
	TcpClinetTcpAcceptCallBindingMethod = t.ClinetBindingMethod
	TcpClinetTcpAcceptCallBindingStruct = "*" + t.ClinetBindingStruct
	Start()
	fmt.Println("TcpStart start")
}
