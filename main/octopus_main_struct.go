package main

import (
	"octopus/log"
)

type MainService struct {
	MainDao *MainDao `autoInjectLang:"main.MainDao"`
	Test    string   `autoInjectCfg:"octopus.director.name"`
}

func (my *MainService) AddMain() {
	log.Info("MainService.AddMain()")
	my.MainDao.AddMain()
}

type MainDao struct {
	MainService *MainService `autoInjectLang:"main.MainService"`
}

func (my *MainDao) AddMain() {
	log.Info("MainDao.AddMain()")
}

type MainUdp struct {
}

func (udp *MainUdp) Call(in interface{}) {
	log.Info("MainUdp=======> ", in)
}

type MainTcpClinet struct {
}

func (tcp *MainTcpClinet) Call(in interface{}) {
	log.Info("MainTcpClinet=======> ", in)
}

type MainTcpServer struct {
}

func (tcp *MainTcpServer) Call(in interface{}) {
	log.Info("MainTcpServer=======> ", in)
}

type MainLogDebugServer struct {
}

func (tcp *MainLogDebugServer) Call(in interface{}) {
	log.Info("MainLogDebugServer=======> ", in)
}

type MainLogInfoServer struct {
}

func (tcp *MainLogInfoServer) Call(in interface{}) {
	log.Info("MainLogInfoServer=======> ", in)
}

type MainLogWarnServer struct {
}

func (tcp *MainLogWarnServer) Call(in interface{}) {
	log.Info("MainLogWarnServer=======> ", in)
}

type MainLogErrorServer struct {
}

func (tcp *MainLogErrorServer) Call(in interface{}) {
	log.Info("MainLogErrorServer=======> ", in)
}
