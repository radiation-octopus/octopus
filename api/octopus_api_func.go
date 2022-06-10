package api

import (
	"sync"
	"time"
)

//import "sync"

var octopusApi *OctopusApi

var once sync.Once

//单例模式
func getInstance() *OctopusApi {
	once.Do(func() {
		octopusApi = new(OctopusApi)
	})
	return octopusApi
}

func Start() {
	getInstance().start()
}

func Stop() {
	getInstance().stop()
}

//绑定api
func BindingApi(
	baseApi BaseApi,
	f func(map[string]interface{}) map[string]interface{},
	requestMethod string,
	path ...string) {
	getInstance().bindingApi(baseApi, f, requestMethod, path)
}

func AddApi(apis ...BaseApi) {
	getInstance().addApi(apis)
}

func CleanSessionSchedule() {
	if IsSession {
		go func() {
			for {
				getInstance().cleanSession()
				time := time.NewTimer(time.Second * 30)
				<-time.C
			}
		}()
	}
}
