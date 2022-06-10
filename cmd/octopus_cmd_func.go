package cmd

import "sync"

var octopusCmd *OctopusCmd

var once sync.Once

//单例模式
func getInstance() *OctopusCmd {
	once.Do(func() {
		octopusCmd = new(OctopusCmd)
	})
	return octopusCmd
}

//初始化cmd
func init() {
	getInstance().cmdInit()
}

//绑定cmd
func BindingCmd(
	baseCmd interface{},
	funcName string,
	cmdName string,
	flagTypeMap map[string]interface{}) {
	getInstance().bindingCmd(baseCmd, funcName, flagTypeMap, cmdName)
}

//执行cmd
func ExecuteCmd() {
	getInstance().executeCmd()
}
