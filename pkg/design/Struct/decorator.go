package Struct

import "fmt"

// 装饰器本身就是要为新的对象创建一些新的行为, 这个行为应该是要服务所有的实例的, 所以要在内部使用抽象接口
// 这里体现的是开闭原则和依赖倒置原则, 装饰器内部依赖的是抽象而不是具体的实现

// 水果抽象基类, 具体的实际类, 然后装饰器要为抽象去做实现!

type FriutBase interface {
	getPrice() int
}

type Apple struct {
}

func (a *Apple) getPrice() int {
	fmt.Println("origin Apple getPrice: 15")
	return 15
}

type DaiziDecorator struct {
	f FriutBase
}

func (d *DaiziDecorator) getPrice() int {
	fmt.Println("add daizi getPrice +3")
	return d.f.getPrice()
}

func RunDecorator() {
	a := &Apple{}
	daizi := &DaiziDecorator{f: a}

	fmt.Println(daizi.getPrice())
}
