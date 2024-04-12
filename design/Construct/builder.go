package Construct

import "fmt"

// 生成器模式是有多种不同的实例，每种实例生成的过程都是相同的，此时用生成器模式可以让每个执行者都具有同样的参数!

type House struct {
	windowType int
	doorType   int
}

type BuilderBase interface { // 这个是用于
	setWin()
	setDoor()

	getHouse() House
}

type NormalBuilder struct {
	windowType int
	doorType   int
}

func (n *NormalBuilder) setWin() {
	n.windowType = 0
}

func (n *NormalBuilder) setDoor() {
	n.doorType = 0
}

func (n *NormalBuilder) getHouse() House {
	return House{
		windowType: n.windowType,
		doorType:   n.doorType,
	}
}

func RunBuilder() {
	bui := &NormalBuilder{}
	bui.setDoor()
	bui.setWin()

	h := bui.getHouse()
	fmt.Println("res: ", h.windowType, " <=> ", h.doorType)
}
