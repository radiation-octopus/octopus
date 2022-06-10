package console

import (
	"bufio"
	"fmt"
	"github.com/radiation-octopus/octopus/core"
	"github.com/radiation-octopus/octopus/director"
	"github.com/radiation-octopus/octopus/utils"
	"os"
	"strings"
)

type Console struct {
	LangName   string
	funcName   string
	formMap    map[string]interface{}
	info       string
	consoleMap map[string]*Console
}

//添加控制
func (c *Console) addConsole(
	LangName string,
	funName string,
	formMap map[string]interface{},
	info string,
	consoleName []string) {
	if c.consoleMap == nil {
		c.consoleMap = make(map[string]*Console)
	}
	consoleMap := c.consoleMap
	for i, s := range consoleName {
		if consoleMap[s] != nil {
			consoleMap = consoleMap[s].consoleMap
		} else if i+1 < len(consoleName) {
			console := new(Console)
			console.consoleMap = make(map[string]*Console)
			consoleMap[s] = console
			consoleMap = consoleMap[s].consoleMap
		} else if i+1 == len(consoleName) {
			console := new(Console)
			console.info = info
			console.formMap = formMap
			console.LangName = LangName
			console.funcName = funName
			console.consoleMap = make(map[string]*Console)
			consoleMap[s] = console
		}
	}
}

//添加控制
func (c *Console) addConsoleAllPath(
	console *Console,
	consoleName []string) {
	if c.consoleMap == nil {
		c.consoleMap = make(map[string]*Console)
	}
	cons := c
	for i, s := range consoleName {
		if strings.Contains(s, "\\$") {
			name := strings.Split(s, ConsolePathAll)[1]
			cons.consoleMap = cons.addConsoleMapAll(name, console)
		} else if s == "" {
			name := strings.Split(consoleName[i+1], ConsolePathAll)[1]
			cons.consoleMap = cons.addConsoleMapAll(name, console)
		}
	}
}

func (c *Console) addConsoleMapAll(name string, addConsole *Console) map[string]*Console {
	consoleMap := c.consoleMap
	for _, v := range consoleMap {
		v.consoleMap = v.addConsoleMapAll(name, addConsole)
	}
	if consoleMap != nil {
		consoleMap[name] = addConsole
	}
	return consoleMap
}

//获取返回Console
func (c *Console) getConsole(consoleName []string) *Console {
	console := c
	if consoleName == nil {
		return console
	}
	for i, s := range consoleName {
		if console.consoleMap[s] == nil {
			return nil
		} else if i+1 <= len(consoleName) {
			console = console.consoleMap[s]
		}
	}
	return console
}

type OctopusConsole struct {
	console            *Console            //控制台map
	currentConsolePath []string            //当前路径
	isExecute          bool                //是否执行
	lastConsole        map[string]*Console //初始化
}

//初始化
func (c *OctopusConsole) consoleInit() {
	c.lastConsole = make(map[string]*Console)
	c.currentConsolePath = []string{}
	c.console = new(Console)
	c.isExecute = true
}

//绑定Console method方法
func (c *OctopusConsole) bindingConsole(
	baseConsole interface{},
	funName string,
	formMap map[string]interface{},
	info string,
	consoleName []string) {
	//注入容器
	director.Register(baseConsole)

	consolepath := strings.Join(consoleName, ".")
	baseConsoleName := core.GetLangName(baseConsole)

	if strings.Contains(consolepath, ConsolePathAll) {
		console := new(Console)
		console.funcName = funName
		console.LangName = baseConsoleName
		console.formMap = formMap
		console.info = info
		c.lastConsole[consolepath] = console
	} else {
		funTypeStr := core.GetLangFuncType(baseConsoleName, funName)
		if funTypeStr != bindingConsoleType {
			panic(bindingConsoleError)
		}

		c.console.addConsole(
			baseConsoleName,
			funName,
			formMap,
			info,
			consoleName)
	}
}

//绑定lastConsole
func (c *OctopusConsole) lastConsoleAdd() {
	for k, v := range c.lastConsole {
		consoleName := strings.Split(k, ".")
		c.console.addConsoleAllPath(v, consoleName)
	}
}

//控制执行
func (c *OctopusConsole) executeConsole() {
	for c.isExecute {
		//请输入命令
		fmt.Println("Please enter the command：")
		reader := bufio.NewReader(os.Stdin)
		//获取输入命令
		consolePath, _ := reader.ReadString('\n')
		consolePath = strings.TrimSuffix(consolePath, "\n")
		//添加命令路径
		c.currentConsolePath = append(c.currentConsolePath, consolePath)
		console := c.console.getConsole(c.currentConsolePath)
		if console == nil {
			c.currentConsolePath = c.currentConsolePath[:len(c.currentConsolePath)-1]
			continue
		}
		//判断是否包含方法
		if console.LangName != "" || console.funcName != "" {
			var isExecuteCall = true
			inMap := make(map[string]interface{})
			if console.formMap != nil {
				//数据按钮
				for k, v := range console.formMap {
					fmt.Println("Please enter the ", k)
					value := bufio.NewReader(os.Stdin)
					//获取输入命令
					form, _ := value.ReadString('\n')
					i := utils.AssignInterface(v, form)
					inMap[k] = i
					fmt.Println("The input values: ", i)
				}
				//数据显示
				fmt.Println("data counter")
				for k, v := range inMap {
					fmt.Println(k, ":", v)
				}
				//确认按钮
				if Isconfirm {
					var cycle = true
					for cycle {
						fmt.Println("Please enter confirmation ", ConfirmConfirm, " or ", CancelConfirm)
						value := bufio.NewReader(os.Stdin)
						toConfirm, _ := value.ReadString('\n')
						if toConfirm == ConfirmConfirm {
							cycle = false
							fmt.Println("Confirm completed")
						} else if toConfirm == CancelConfirm {
							isExecuteCall = false
							cycle = false
							fmt.Println("Cancel the confirmation")
						} else {
							continue
						}
					}
				}
			}
			//是否确认收到
			if isExecuteCall {
				//执行命令方法回调
				out := core.CallMethod(console.LangName, console.funcName, inMap)
				fmt.Println("Return after running the program：", out)
			}
		}
	}
}
