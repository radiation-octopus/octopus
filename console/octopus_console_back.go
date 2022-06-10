package console

type BackConsole struct {
}

//返回上一级方法
func (c *BackConsole) BackCmd(inMap map[string]interface{}) interface{} {
	currentConsolePath := getInstance().currentConsolePath
	if len(currentConsolePath) > 1 {
		currentConsolePath = currentConsolePath[:len(currentConsolePath)-2]
		getInstance().currentConsolePath = currentConsolePath
	} else {
		currentConsolePath = []string{}
		getInstance().currentConsolePath = currentConsolePath
	}
	return true
}
