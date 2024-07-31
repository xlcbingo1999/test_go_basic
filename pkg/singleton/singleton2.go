package singleton

import (
	"log"
	"sync"
	"time"
)

type SharedInstance struct {
	num int
}

// 并发单例 用sync.Once{}
var sharedInstance *SharedInstance

func RunSingleton2() {
	once := &sync.Once{}

	for i := 0; i < 10; i++ {
		go func(index int) {
			once.Do(func() { // 保证只会执行一次
				sharedInstance = &SharedInstance{num: index}
				log.Println("set sharedInstance: ", sharedInstance)
			})
		}(i)
	}

	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		go func(ins *SharedInstance) {
			log.Println("sharedInstance: ", ins)
		}(sharedInstance)
	}
	time.Sleep(2 * time.Second)
}
