package Struct

import "fmt"

// 主要面对的场景是笛卡尔积, 一种设备要和另外的N种设备进行交互的时使用
// Adapter模式： Adapter似乎只是针对一种场景, 某种设备要插入到一种具体的机器上, 需要专门做一个适配器
// Bridge模式: 更加通用的场景, 更加抽象一层, 需要set函数进行动态地修改

type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Printer interface {
	PrintFile()
}

type ComA struct {
	printer Printer
}

func (a *ComA) Print() {
	a.printer.PrintFile()
}

func (a *ComA) SetPrinter(p Printer) {
	a.printer = p
}

type HP struct{}

func (h *HP) PrintFile() {
	fmt.Println("HP printFile")
}

func RunBridge() {
	h := &HP{}
	comA := &ComA{printer: h}

	comA.Print()

}
