package Struct

import "fmt"

type Machine interface {
	DoMethod()
}

type Mac struct {
}

func (m *Mac) DoMethod() {
	fmt.Println("Mac InsertMethod")
}

type Window struct {
}

func (w *Window) DoMethodCuda() {
	fmt.Println("Window InsertMethodCuda")
}

type WindowAdapter struct {
	windowMachine *Window
}

func (w *WindowAdapter) DoMethod() { // 实现的是通用的machine的方法
	fmt.Println("update input")
	w.windowMachine.DoMethodCuda()
	fmt.Println("update output")
}

type Client struct {
}

func (c *Client) DoMethod(machine Machine) {
	machine.DoMethod()
}

func RunAdapter() {
	cli := &Client{}
	mac := &Mac{}

	cli.DoMethod(mac)

	window := &Window{}
	windowApa := &WindowAdapter{windowMachine: window}
	cli.DoMethod(windowApa)
}
