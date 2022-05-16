package main

import (
	"fmt"
	"octopus/core"
	"octopus/director"
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
	testMainService := core.GetLang("main.MainService").(*MainService)
	fmt.Println(testMainDao, testMainService)
}
