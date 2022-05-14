package director

import "fmt"

//DirectorStart
type DirectorStart struct {
}

func (d *DirectorStart) Start() {
	fmt.Println("DirectorStart start")
}
