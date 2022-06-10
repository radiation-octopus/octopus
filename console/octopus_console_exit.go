package console

type ExitConsole struct {
}

//退出方法
func (c *ExitConsole) ExitCmd(inMap map[string]interface{}) interface{} {
	getInstance().isExecute = false
	return true
}
