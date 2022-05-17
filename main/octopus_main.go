package main

import (
	_ "octopus/blockchain"
	"octopus/core"
	"octopus/director"
	"octopus/log"
	_ "octopus/log"
	_ "octopus/tcp"
	_ "octopus/udp"
)

func init() {
	mainDao := new(MainDao)
	mainService := new(MainService)
	mainUdp := new(MainUdp)
	mainTcpServer := new(MainTcpServer)
	mainTcpClinet := new(MainTcpClinet)
	director.Register(mainDao)
	director.Register(mainService)
	director.Register(mainUdp)
	director.Register(mainTcpServer)
	director.Register(mainTcpClinet)
}

func main() {
	director.Start()
	var testMainDao = core.GetLang("main.MainDao").(*MainDao)
	var testMainService = core.GetLang("main.MainService").(*MainService)
	log.Info(testMainDao, testMainService)
}
