package console

import "sync"

var octopusConsole *OctopusConsole

var once sync.Once

//单例模式
func getInstance() *OctopusConsole {
	once.Do(func() {
		octopusConsole = new(OctopusConsole)
	})
	return octopusConsole
}

//绑定console
func BindingConsole(
	baseConsole interface{},
	funcName string,
	flagTypeMap map[string]interface{},
	info string,
	consoleName ...string) {
	getInstance().bindingConsole(baseConsole, funcName, flagTypeMap, info, consoleName)
}

//执行console
func ExecuteConsole() {
	getInstance().executeConsole()
}
