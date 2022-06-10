package gorm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type OctopusGorm struct {
	db    *gorm.DB //gorm链接mysql
	sqldb *sql.DB
}

func (m *OctopusGorm) start() {
	//:= DirverName
	host := Host
	port := strconv.Itoa(Port)
	database := Database
	username := Username
	password := Password
	charset := Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	m.db, _ = gorm.Open(mysql.Open(args), &gorm.Config{})
	m.sqldb, _ = m.db.DB()
	m.sqldb.SetMaxIdleConns(PoolNumActive)
	m.sqldb.SetMaxOpenConns(PoolNumMax)
	defer m.sqldb.Close()
}

func (m *OctopusGorm) stop() {
}

func (m *OctopusGorm) getDb() *gorm.DB {
	return m.db
}
