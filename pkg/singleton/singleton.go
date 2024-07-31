package singleton

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Singleton struct {
	val int
}

var instance *Singleton
var initialized uint32 // 原子操作的判断只有Uint32是最小的了
var mu sync.Mutex
var once sync.Once // Go 原生提供的单例支持

func GetNoThreadInstance(_val int) *Singleton { // 线程不安全的方法
	if instance == nil {
		instance = &Singleton{val: _val}
	}

	return instance
}

func GetThreadInstance(_val int) *Singleton {
	// 每次都要加锁, 其实如果存在的时候直接从缓存读取即可!
	for !mu.TryLock() {
		runtime.Gosched() // 让给别的routine先进行操作, 但这个让的时间是一个内部设置好的时间, 不会一直都让出去的
	}

	defer mu.Unlock()

	if instance == nil {
		instance = &Singleton{val: _val}
	}
	return instance
}

func GetCheckLockCheckThreadInstance(_val int) *Singleton {
	if instance == nil { // check, 这里不是完全原子的! 此时可能有多个线程都走进了 instance == nil的逻辑中
		for !mu.TryLock() { // lock
			runtime.Gosched()
		}

		defer mu.Unlock()

		if instance == nil { // check, 避免别的线程已经修改了这个值, 但我们刚刚从锁里面出来, 需要重新判断以下
			instance = &Singleton{val: _val}
		}
	}
	return instance
}

func GetAtomicThreadInstance(_val int) *Singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	for !mu.TryLock() {
		runtime.Gosched()
	}
	defer mu.Unlock()

	if initialized == 0 {
		instance = &Singleton{val: _val}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

func GetOnceThreadInstance(_val int) *Singleton {
	once.Do(func() {
		instance = &Singleton{val: _val}
	}) // 在内部判断once中的原子量o.done是否被设置为1, 如果为0则调用传进去的函数, 完成函数后就可以设置o.done为1, 此时非常完美地实现任何的单例的构造
	return instance
}

func RunTest() {
	// 1. 不安全读写
	testCase := 20
	for i := 0; i < testCase; i++ {
		go func(index int) {
			resInstance := GetNoThreadInstance(index)
			fmt.Printf("index: %d => resInstance: %d\n", index, resInstance.val)
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Printf("\n\n\n\n")

	// 2. 完全加锁读写
	for i := 0; i < testCase; i++ {
		go func(index int) {
			resInstance := GetThreadInstance(index)
			fmt.Printf("index: %d => resInstance: %d\n", index, resInstance.val)
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Printf("\n\n\n\n")

	// 3. check-lock-check读写
	for i := 0; i < testCase; i++ {
		go func(index int) {
			resInstance := GetCheckLockCheckThreadInstance(index)
			fmt.Printf("index: %d => resInstance: %d\n", index, resInstance.val)
		}(i)
	}

	time.Sleep(10 * time.Second)
	fmt.Printf("\n\n\n\n")

	// 4. atomic读写
	for i := 0; i < testCase; i++ {
		go func(index int) {
			resInstance := GetAtomicThreadInstance(index)
			fmt.Printf("index: %d => resInstance: %d\n", index, resInstance.val)
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Printf("\n\n\n\n")

	// 5. sync.Once读写
	for i := 0; i < testCase; i++ {
		go func(index int) {
			resInstance := GetOnceThreadInstance(index)
			fmt.Printf("index: %d => resInstance: %d\n", index, resInstance.val)
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Printf("\n\n\n\n")
}
