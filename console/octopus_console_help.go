package console

type HelpConsole struct {
}

type helpReturn struct {
	command string
	info    string
}

//查询命令帮助
func (c *HelpConsole) HelpCmd(inMap map[string]interface{}) interface{} {
	currentConsolePath := getInstance().currentConsolePath
	if len(currentConsolePath) >= 1 {
		currentConsolePath = currentConsolePath[:len(currentConsolePath)-1]
		getInstance().currentConsolePath = currentConsolePath
	}
	console := getInstance().console
	consoleMap := console.getConsole(currentConsolePath).consoleMap
	helpReturns := []helpReturn{}
	for k, console := range consoleMap {
		helpReturn := new(helpReturn)
		helpReturn.command = k
		helpReturn.info = console.info
		helpReturns = append(helpReturns, *helpReturn)
	}
	return helpReturns
}
