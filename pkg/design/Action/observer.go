package Action

import "fmt"

type Observer interface {
	doAction()
}

type CustomObserver struct {
	id int
}

func (co *CustomObserver) doAction() {
	fmt.Println("id: ", co.id, " doAction")
}

type ObserveSubject interface {
	register(Observer)
	deregister(Observer)
	notifyAll()
}

type ObItem struct {
	name  string
	state bool

	observers []Observer
}

func (i *ObItem) updateState() {
	i.state = !i.state
	i.notifyAll()
}

func (i *ObItem) register(ob Observer) {
	for _, cob := range i.observers {
		if cob == ob {
			return
		}
	}
	i.observers = append(i.observers, ob)
}

func (i *ObItem) deregister(ob Observer) {
	targetIndex := -1
	for index, cob := range i.observers {
		if cob == ob {
			targetIndex = index
			break
		}
	}
	if targetIndex == -1 {
		return
	} else {
		i.observers = append(i.observers[0:targetIndex], i.observers[targetIndex+1:]...)
	}
}

func (i *ObItem) notifyAll() {
	for _, cob := range i.observers {
		cob.doAction()
	}
}

func RunObserver() {
	item := &ObItem{name: "shirt", state: false}

	cus1 := &CustomObserver{id: 1}
	cus2 := &CustomObserver{id: 2}

	item.register(cus1)
	item.register(cus2)

	item.updateState()
}
