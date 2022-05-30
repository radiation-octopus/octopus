package db

import "github.com/radiation-octopus/octopus/director"

//初始化octopus
func init() {
	//把启动注入
	director.Register(new(DbStart))
	//把停止注入
	director.Register(new(DbStop))
}
