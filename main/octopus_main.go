package main

import (
	_ "github.com/radiation-octopus/octopus/api"
	"github.com/radiation-octopus/octopus/console"
	_ "github.com/radiation-octopus/octopus/console"
	_ "github.com/radiation-octopus/octopus/db"
	"github.com/radiation-octopus/octopus/director"
	_ "github.com/radiation-octopus/octopus/gorm"
	_ "github.com/radiation-octopus/octopus/log"
	_ "github.com/radiation-octopus/octopus/tcp"
	_ "github.com/radiation-octopus/octopus/udp"
)

func main() {
	director.Start()
	console.ExecuteConsole()
	//fmt.Println(reflect.TypeOf(ceshi))
	//var k = reflect.TypeOf(0)
	//var e = reflect.TypeOf("")
	//fmt.Println(reflect.FuncOf([]reflect.Type{k},
	//	[]reflect.Type{e}, false).String())
}
