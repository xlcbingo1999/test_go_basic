package preemptX

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("This is f1")
}

func f2() {
	fmt.Println("This is f2")
}

func f3() {
	// 死循环逻辑
	for {

	}
	fmt.Println("This is f3")
}

func RunPreemptTest1() {
	// 保证此时只有一个CPU
	runtime.GOMAXPROCS(1)

	// 创建G1, 会放置在本地队列的头部
	go f1()

	// 创建G2, 放置在本地队列的头部
	go f2()

	// 等待 f1 f2的执行 gopark 主 goroutine GMP调度可运行的 G, 按顺序调用 f2 f1
	time.Sleep(100 * time.Millisecond)
	fmt.Println("success")
}

func RunPreemptTest2() {
	// 保证此时只有一个CPU
	runtime.GOMAXPROCS(1)

	// 创建G1, 会放置在本地队列的头部
	go f1()

	// 创建G2, 放置在本地队列的头部
	go f2()

	// 创建死循环的方法
	go f3()

	// 等待 f1 f2的执行 gopark 主 goroutine GMP调度可运行的 G, 按顺序调用 f3 f2 f1
	// 但实际上, golang实现了抢占式调度, 可以避免f3()一直执行, f2和f1也会输出
	time.Sleep(100 * time.Millisecond)
	fmt.Println("success")
}

func RunPreempt() {
	RunPreemptTest2()
}
