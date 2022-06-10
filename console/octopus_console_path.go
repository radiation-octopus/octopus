package console

type PathConsole struct {
}

//获得执行路径
func (c *PathConsole) PathCmd(inMap map[string]interface{}) interface{} {
	currentConsolePath := getInstance().currentConsolePath
	if len(currentConsolePath) >= 1 {
		currentConsolePath = currentConsolePath[:len(currentConsolePath)-1]
		getInstance().currentConsolePath = currentConsolePath
	}
	return currentConsolePath
}
