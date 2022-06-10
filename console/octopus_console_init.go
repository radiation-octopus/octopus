package console

import "github.com/radiation-octopus/octopus/director"

func init() {
	getInstance().consoleInit()
	director.Register(new(ConsoleStart))
	director.Register(new(ConsoleStop))
	BindingConsole(
		new(BackConsole),
		"BackCmd",
		nil,
		"Return to the previous level",
		"\\$back")
	BindingConsole(
		new(ExitConsole),
		"ExitCmd",
		nil,
		"Exit all programs",
		"\\$exit")
	BindingConsole(
		new(HelpConsole),
		"HelpCmd",
		nil,
		"Help view details",
		"\\$help")
	BindingConsole(
		new(PathConsole),
		"PathCmd",
		nil,
		"Obtaining the execution path",
		"\\$path")

	BindingConsole(
		new(TestConsole),
		"TestCmd",
		nil,
		"test",
		"test")

	BindingConsole(
		new(Test1Console),
		"Test1Cmd",
		nil,
		"test",
		"test", "test1")
}

type TestConsole struct {
}

func (c *TestConsole) TestCmd(inMap map[string]interface{}) interface{} {
	return true
}

type Test1Console struct {
}

func (c *Test1Console) Test1Cmd(inMap map[string]interface{}) interface{} {
	return true
}
