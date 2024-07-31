package futuresX

import (
	"log"
	"time"
)

type Matrix struct{}

func Inverse(_ *Matrix) interface{} {
	time.Sleep(1 * time.Second)
	return 0
}

func InverseFuture(m *Matrix) chan interface{} {
	result := make(chan interface{})

	go func() {
		result <- Inverse(m)
	}()

	return result
}

func Product(a interface{}, b interface{}) int {
	return 4
}

func ProductFuture(a chan interface{}, b chan interface{}) chan int {
	result := make(chan int)
	go func(result chan int) {
		realA := <-a
		log.Println("finished load a")
		realB := <-b
		log.Println("finished load b")

		result <- Product(realA, realB)
	}(result)

	return result
}

func RunFuture() {
	a := &Matrix{}
	b := &Matrix{}
	aInvFuture := InverseFuture(a)
	bInvFuture := InverseFuture(b)
	resultFuture := ProductFuture(aInvFuture, bInvFuture)
	log.Println("define all logic")

	for {
		select {
		case result := <-resultFuture:
			log.Println("get final result: ", result) // 管道是会阻塞等待的, 所以是需要额外的同步逻辑的, 可以选择select让其执行其他逻辑
			log.Println("close all")
			return
		default:
			log.Println("can do other things!")
			time.Sleep(100 * time.Millisecond)
		}
	}

}
