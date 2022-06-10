package gorm

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"sync"
)

var octopusGorm *OctopusGorm

var once sync.Once

//单例模式
func getInstance() *OctopusGorm {
	once.Do(func() {
		octopusGorm = new(OctopusGorm)
	})
	return octopusGorm
}

func Start() {
	getInstance().start()
}

func Stop() {
	getInstance().stop()
}

func GetDb() *gorm.DB {
	return getInstance().getDb()
}
