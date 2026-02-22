package cmdmanager

import "fmt"

type CMDManager struct {
	
	
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("enter")
}

func (cmd CMDManager) WriteResult() (data interface{}) {
	fmt.Println(data)
	return nil
}