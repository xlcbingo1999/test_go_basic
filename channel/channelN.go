package channel

import (
	"fmt"
	"sync"
)

func printCir(inCh chan struct{}, outCh chan struct{}, index int, N int, fin int, wg *sync.WaitGroup) {
	defer (*wg).Done()
	val := index
	for {
		if val >= fin {
			break
		}
		if index == 0 {
			fmt.Println("val: ", val)
			val += N
			outCh <- struct{}{}
			<-inCh
		} else {
			<-inCh
			fmt.Println("val: ", val)
			val += N
			outCh <- struct{}{}
		}
	}
}

func RunChannelN() {
	N := 5
	allCh := [](chan struct{}){}
	for i := 0; i < N; i++ {
		allCh = append(allCh, make(chan struct{})) // 每个channel都必须自己手动进行初始化
	}
	fin := 100
	wg := sync.WaitGroup{}
	wg.Add(N)

	for i := 0; i < N-1; i++ {
		go printCir(allCh[i], allCh[i+1], i, N, fin, &wg)
	}
	go printCir(allCh[N-1], allCh[0], N-1, N, fin, &wg)

	wg.Wait()

	for i := 0; i < N; i++ {
		close(allCh[i])
	}
	fmt.Println("all finished")
}
