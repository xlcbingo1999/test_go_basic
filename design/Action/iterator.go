package Action

// 迭代器模式的核心: 用hasNext和getNext来获取下一个资源, 迭代器要把容器包含在内

import "fmt"

type Item struct {
	val int
}

func NewItem(v int) *Item {
	return &Item{
		val: v,
	}
}

type Iterator interface {
	hasNext() bool
	getNext() *Item
}

type RealIterator struct {
	currIndex int
	container []*Item
}

func (ri *RealIterator) hasNext() bool {
	return ri.currIndex < len(ri.container)
}

func (ri *RealIterator) getNext() *Item {
	res := ri.container[ri.currIndex]
	ri.currIndex += 1
	return res
}

func RunIterator() {
	ri := &RealIterator{
		currIndex: 0,
		container: make([]*Item, 4),
	}
	for i := 0; i < 4; i++ {
		ri.container[i] = &Item{val: i}
	}

	for ri.hasNext() {
		next := ri.getNext()
		fmt.Println("next: ", next.val)
	}
}
