package director

import "fmt"

//Director停止方法
type DirectorStop struct {
}

func (d *DirectorStop) Stop() {
	fmt.Println("DirectorStop stop")
}
