package gorm

import "github.com/radiation-octopus/octopus/log"

type GormStop struct {
}

func (m *GormStop) Stop() {
	Stop()
	log.Info("GormStop Stop")
}
