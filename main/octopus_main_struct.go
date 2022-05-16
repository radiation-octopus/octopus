package main

import (
	"fmt"
)

type MainService struct {
	MainDao *MainDao `autoInjectLang:"main.MainDao"`
	Test    string   `autoInjectCfg:"octopus.director.name"`
}

func (my *MainService) AddMain() {
	fmt.Println("MainService.AddMain()")
	my.MainDao.AddMain()
}

type MainDao struct {
	MainService *MainService `autoInjectLang:"main.MainService"`
}

func (my *MainDao) AddMain() {
	fmt.Println("MainDao.AddMain()")
}

type MainUdp struct {
}

func (udp *MainUdp) Call(in interface{}) {
	fmt.Println("MainUdp=======> ", in)
}

type MainTcpClinet struct {
}

func (tcp *MainTcpClinet) Call(in interface{}) {
	fmt.Println("MainTcpClinet=======> ", in)
}

type MainTcpServer struct {
}

func (tcp *MainTcpServer) Call(in interface{}) {
	fmt.Println("MainTcpServer=======> ", in)
}
