package cmd

import (
	"flag"
	"github.com/radiation-octopus/octopus/core"
	"github.com/radiation-octopus/octopus/director"
	"github.com/radiation-octopus/octopus/log"
	"os"
	"reflect"
)

type Cmd struct {
	LangName string
	funcName string
	inMap    map[string]interface{}
	out      bool
}

type OctopusCmd struct {
	CmdMap map[string]*Cmd //路径map
}

func (c *OctopusCmd) cmdInit() {
	c.CmdMap = make(map[string]*Cmd)
}

//绑定Cmd method方法
func (c *OctopusCmd) bindingCmd(
	baseCmd interface{},
	funName string,
	flagTypeMap map[string]interface{},
	cmdName string) {
	inMap := make(map[string]interface{})
	cmdNameCmd := flag.NewFlagSet(cmdName, flag.ExitOnError)
	for k, v := range flagTypeMap {
		vType := reflect.TypeOf(&v).String()
		vValue := reflect.ValueOf(&v)
		switch vType {
		case "*string":
			inMap[k] = *cmdNameCmd.String(k, vValue.String(), k)
		case "*int":
			inMap[k] = *cmdNameCmd.Int(k, int(vValue.Int()), k)
		case "*float64":
			inMap[k] = *cmdNameCmd.Float64(k, vValue.Float(), k)
		case "*bool":
			inMap[k] = *cmdNameCmd.Bool(k, vValue.Bool(), k)
		}
	}
	director.Register(baseCmd)
	baseCmdName := reflect.TypeOf(baseCmd).String()
	funTypeStr := core.GetLangFuncType(baseCmdName, funName)
	if funTypeStr != bindingCmdType {
		panic(bindingCmdError)
	}
	cmd := new(Cmd)
	cmd.inMap = inMap
	cmd.LangName = baseCmdName
	cmd.funcName = funName
	c.CmdMap[cmdName] = cmd
}

func (c *OctopusCmd) executeCmd() {
	// 期望前面定义的子命令作为第一个参数传入。
	if len(os.Args) < 2 {
		panic(executeCmdError)
		os.Exit(1)
	}
	cmd := c.CmdMap[os.Args[1]]
	if cmd != nil {
		cmd.out = core.CallMethod(cmd.LangName, cmd.funcName, cmd.inMap)[0].(bool)
	} else {
		panic(executeCmdError)
	}
	log.Info(cmd.out)
}
