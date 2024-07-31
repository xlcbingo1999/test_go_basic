package Action

import "fmt"

// 责任链模式的核心就是next()函数, 每个部门的下一个部门是耦合确定的,
// 每个部门实例都要维护一个next部门, 执行的时候会自动执行下一个部门的方法

type Patient struct {
	finishedA bool
	finishedB bool
}

type Department interface {
	execute(*Patient)
	SetNext(Department)
}

type ADe struct {
	next Department
}

func (a *ADe) execute(p *Patient) {
	if p.finishedA {
		fmt.Println("has finished A")
		a.next.execute(p)
	}
	fmt.Println("execute A")
	p.finishedA = true
	a.next.execute(p)
}

func (a *ADe) SetNext(d Department) {
	a.next = d
}

type BDe struct { // 最后一个业务, 不再需要next
	next Department
}

func (b *BDe) execute(p *Patient) {
	if p.finishedB {
		fmt.Println("has finished B")
		// b.next.execute(p)
	}
	fmt.Println("execute B")
	p.finishedB = true
	// b.next.execute(p)
}

func (b *BDe) SetNext(d Department) {
	b.next = d
}

func RunResponsibility() {
	p := &Patient{
		finishedA: false,
		finishedB: false,
	}

	a := &ADe{}
	b := &BDe{}
	a.SetNext(b)
	a.execute(p)
}
