package channel

import (
	"fmt"
	"sync"
)

func print1(ch1 chan struct{}, ch2 chan struct{}, wg *sync.WaitGroup) {
	defer (*wg).Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("%d: %d\n", 1, i)
		ch2 <- struct{}{}
		<-ch1
	}
}

func print2(ch2 chan struct{}, ch3 chan struct{}, wg *sync.WaitGroup) {
	defer (*wg).Done()
	for i := 0; i < 100; i++ {
		<-ch2
		fmt.Printf("%d: %d\n", 2, i)
		ch3 <- struct{}{}
	}
}

func print3(ch3 chan struct{}, ch1 chan struct{}, wg *sync.WaitGroup) {
	defer (*wg).Done()
	for i := 0; i < 100; i++ {
		<-ch3
		fmt.Printf("%d: %d\n", 3, i)
		ch1 <- struct{}{}
	}
}

func Run123() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	go print1(ch1, ch2, &wg)
	go print2(ch2, ch3, &wg)
	go print3(ch3, ch1, &wg)

	wg.Wait() // wait在这里阻塞住, 此时没法走到逻辑内
	fmt.Println("\nwg finished")

	// <-ctx.Done()
	// fmt.Println("all finished")
	close(ch1) // 关闭channel的时候, 当前channel的读写会被处理掉
	close(ch2)
	close(ch3)
}
