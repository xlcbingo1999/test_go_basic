package singleton

import "sync"

// 1. 饿汉式: 在func init()中就定义单例实例出来
// 2. 懒汉式: 在GetInstance()中，判断单例实例是否存在，不存在的时候再进行实例化
// 3. 懒汉式+并发安全式: GetInstance()的过程中需要加锁和解锁
// 4. check-lock-check式(双重锁定式): 避免每次获取单例都要加锁
// 5. once.Do()方式: 内部通过atomic操作, 进行Compare-and-Swap

type MySingleton struct{}

var mymu sync.Mutex
var single *MySingleton

func GetSingleton3() *MySingleton {
	mymu.Lock()
	defer mymu.Unlock()

	if single == nil {
		single = &MySingleton{}
	}
	return single
}

func GetSingletonCheckLockCheck() *MySingleton {
	if single == nil {
		mymu.Lock()
		defer mu.Unlock()

		if single == nil { // 极端情况下多个线程会走到这一步, 需要做一个预备处理
			single = &MySingleton{}
		}
	}
	return single
}

var myOnce sync.Once // 内部是atomic操作

func GetSingletonOnce() *MySingleton {
	if single == nil {
		myOnce.Do(func() {
			single = &MySingleton{}
		})
	}

	return single
}

func RunMySingleton() {

}
