package design

import (
	"fmt"

	"github.com/xlcbingo1999/test_go_basic/design/Action"
	"github.com/xlcbingo1999/test_go_basic/design/Construct"
	"github.com/xlcbingo1999/test_go_basic/design/Struct"
)

func RunAllDesign() {
	Construct.RunAbstractFactory()
	fmt.Println("==============")

	Construct.RunBuilder()
	fmt.Println("==============")

	Construct.RunProtocal()
	fmt.Println("==============")

	Struct.RunAdapter()
	fmt.Println("==============")

	Struct.RunBridge()
	fmt.Println("==============")

	Struct.RunCombine()
	fmt.Println("==============")

	Struct.RunDecorator()
	fmt.Println("==============")

	Struct.RunProxy()
	fmt.Println("==============")

	Action.RunResponsibility()
	fmt.Println("==============")

	Action.RunCommand()
	fmt.Println("==============")

	Action.RunIterator()
	fmt.Println("==============")

	Action.RunMediator()
	fmt.Println("==============")

	Action.RunObserver()
	fmt.Println("==============")

	Action.RunState()
	fmt.Println("==============")

	Action.RunStrategy()
	fmt.Println("==============")
}
