package tcp

import "fmt"

//Tcp停止方法
type TcpStop struct {
}

func (t *TcpStart) Stop() {
	Stop()
	fmt.Println("TcpStop stop")
}
