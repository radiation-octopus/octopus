package console

import "github.com/radiation-octopus/octopus/log"

type ConsoleStart struct {
	Isconfirm      bool   `autoInjectCfg:"octopus.console.confirm.is"`
	ConfirmConfirm string `autoInjectCfg:"octopus.console.confirm.confirm"`
	CancelConfirm  string `autoInjectCfg:"octopus.console.confirm.cancel"`
	IsUseHelp      bool   `autoInjectCfg:"octopus.console.use.help.is"`
	IsUseExit      bool   `autoInjectCfg:"octopus.console.use.exit.is"`
	IsUseBack      bool   `autoInjectCfg:"octopus.console.use.back.is"`
	PathAll        string `autoInjectCfg:"octopus.console.path.all"`
}

//console 启动
func (c *ConsoleStart) Start() {
	Isconfirm = c.Isconfirm
	ConfirmConfirm = c.ConfirmConfirm
	CancelConfirm = c.CancelConfirm
	IsUseHelp = c.IsUseHelp
	IsUseExit = c.IsUseExit
	IsUseBack = c.IsUseBack
	ConsolePathAll = c.PathAll
	getInstance().lastConsoleAdd()
	log.Info("ConsoleStart start")
}
