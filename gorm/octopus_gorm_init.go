package gorm

import "github.com/radiation-octopus/octopus/director"

func init() {
	director.Register(new(GormStart))
	director.Register(new(GormStop))
}
