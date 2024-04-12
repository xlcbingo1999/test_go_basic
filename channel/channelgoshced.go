package channel

import (
	"fmt"
	"runtime"
	"sync"
)

func print1a(wg *sync.WaitGroup) {
	defer (*wg).Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("1: %d\n", i)
		runtime.Gosched()
	}

}

func print2a(wg *sync.WaitGroup) {
	defer (*wg).Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("2: %d\n", i)
		runtime.Gosched()
	}
}

func RunChannelGoSched() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	runtime.GOMAXPROCS(1) // 设置CPU最大核心数为1

	go print1a(&wg)
	go print2a(&wg)

	wg.Wait()
	fmt.Println("RunChannelGoSched")
}
