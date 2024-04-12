package Action

import "fmt"

// 策略模式的核心: 把具体的容器和容器应用的调度策略进行解耦, 可以随着流程的进行动态修改对应的策略
// 类似K8S中的自定义调度器

type StItem interface{}

type StItemSpec struct {
	val int
}

type StContainer interface {
	add(StItem)
	delete(int) StItem
}

type StContainerSpec struct {
	container []StItem
	strategy  Strategy
}

func (c *StContainerSpec) add(item StItem) {
	c.container = append(c.container, item)
}

func (c *StContainerSpec) delete(index int) StItem {
	item := c.container[index]
	c.container = append(c.container[0:index], c.container[index+1:])
	return item
}

type Strategy interface {
	in(StContainer, StItem)
	out(StContainer) StItem
}

type FIFOStrategy struct {
}

func (f *FIFOStrategy) in(container StContainer, item StItem) {
	fmt.Println("in: ", item)
	container.add(item)
}

func (f *FIFOStrategy) out(container StContainer) StItem {
	item := container.delete(0)
	fmt.Println("out: ", item)
	return item
}

func RunStrategy() {
	fifo := &FIFOStrategy{}
	cache := &StContainerSpec{
		container: make([]StItem, 2),
		strategy:  fifo,
	}

	for i := 0; i < 2; i++ {
		cache.container[i] = &StItemSpec{
			val: i,
		}
	}

	fifo.in(cache, &StItemSpec{val: 2})
	res := fifo.out(cache)
	fmt.Println(res.(*StItemSpec).val)
}
