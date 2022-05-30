package api

import "github.com/radiation-octopus/octopus/log"

//Web停止方法
type WebStop struct {
}

func (w *WebStop) Stop() {
	Stop()
	log.Info("WebStop stop")
}
