package api

import (
	"log"
	"net/http"
	"strconv"
)

type OctopusApiHandle struct {
}

func initHandle() *OctopusApiHandle {
	h := new(OctopusApiHandle)
	http.HandleFunc("/", h.httpHandle)
	httpListen()
	return h
}

//httplisten
func httpListen() {
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(Port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

//http监听回调方法
func (h *OctopusApiHandle) httpHandle(writer http.ResponseWriter, request *http.Request) {
	p := initProcess(writer, request)
	p.executeProcess()
}
