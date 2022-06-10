package gorm

import "github.com/radiation-octopus/octopus/log"

type GormStart struct {
	Dirver        string `autoInjectCfg:"octopus.gorm.dirver"`
	Host          string `autoInjectCfg:"octopus.gorm.host"`
	Port          int    `autoInjectCfg:"octopus.gorm.port"`
	Database      string `autoInjectCfg:"octopus.gorm.database"`
	Username      string `autoInjectCfg:"octopus.gorm.username"`
	Password      string `autoInjectCfg:"octopus.gorm.password"`
	Charset       string `autoInjectCfg:"octopus.gorm.charset"`
	PoolNumActive int    `autoInjectCfg:"octopus.gorm.pool.active"`
	PoolNumMax    int    `autoInjectCfg:"octopus.gorm.pool.max"`
}

func (m *GormStart) Start() {
	Dirver = m.Dirver
	Host = m.Host
	Port = m.Port
	Database = m.Database
	Username = m.Username
	Password = m.Password
	Charset = m.Charset
	PoolNumActive = m.PoolNumActive
	PoolNumMax = m.PoolNumMax
	Start()
	log.Info("GormStart Start")
}
