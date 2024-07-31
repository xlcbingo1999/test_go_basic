package Action

// 中介者模式的核心是: 用一个中介者对状态进行统一的管理, 中介者处理的时候会同时进行相应的行动处理
// 需要一个等待队列用于排队, 排队中的元素出队列的时候就调用相应的方法即可处理
// 用户只需要直接将任务提交上去即可, 不需要考虑内部是如何进行中介者管理的

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrive()
}

type HeXieTrain struct {
	mediator Mediator // 做一个delegate模式?
}

func (hx *HeXieTrain) arrive() {
	if !hx.mediator.canArrive(hx) {
		// 非阻塞返回
		fmt.Println("hx not in , noblocking")
		return
	}
	fmt.Println("hx in")
}

func (hx *HeXieTrain) depart() {
	fmt.Println("hx out")
	hx.mediator.notifyAboutDeparture() // 通知另一个火车可以进来了
}

func (hx *HeXieTrain) permitArrive() {
	fmt.Println("hx permit")
	hx.arrive()
}

type Mediator interface { // 中介方法, 需要设置一些相应的方法用于实现
	canArrive(Train) bool  // 判断其是否可以进来
	notifyAboutDeparture() // 管理方案, 用于让某种火车进来
}

type Manager struct {
	waits  []Train
	isBusy bool
}

func (m *Manager) canArrive(t Train) bool {
	if m.isBusy {
		m.waits = append(m.waits, t)
		return false
	}
	m.isBusy = true
	return true
}

func (m *Manager) notifyAboutDeparture() {
	if m.isBusy {
		m.isBusy = false
	}
	if len(m.waits) > 0 {
		front := m.waits[0]
		m.waits = m.waits[1:]
		front.permitArrive()
	}
}

func RunMediator() {
	man := &Manager{
		waits:  make([]Train, 0),
		isBusy: false,
	}

	a := &HeXieTrain{mediator: man}
	b := &HeXieTrain{mediator: man}

	a.arrive()
	b.arrive()
	a.depart()
}
