package mapx

import (
	"context"
	"log"
	"sync"
	"time"
)

type RWMap struct {
	sync.Mutex
	mp     map[int]int // kv存储
	key2ch map[int]chan struct{}
}

func NewRWMap() *RWMap {
	return &RWMap{
		mp:     make(map[int]int),
		key2ch: make(map[int]chan struct{}),
	}
}

func (m *RWMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()

	m.mp[k] = v
	ch, exist := m.key2ch[k]
	if !exist { // 如果没有相关的管道, 就不需要额外的阻塞逻辑了, 直接结束就可以了
		return
	}
	// m.key2ch[k] <- struct{}{} // 写入之后需要有人进行消耗, 不然会阻塞
	select {
	case <-ch: // 可能本来就有内容
		return
	default:
		close(ch) // 主动去通知所有的channel! 所有阻塞的读写goroutine都会被唤醒, 写goroutine会进入panic, 读goroutine不会
	}
}

func (m *RWMap) Get(k int, maxWaitDuration time.Duration) (int, error) {
	m.Lock()

	v, exist := m.mp[k]
	if exist {
		m.Unlock()
		return v, nil
	}

	// 新建一个管道, 用于等待写入值
	ch, exist := m.key2ch[k]
	if !exist {
		ch = make(chan struct{})
		m.key2ch[k] = ch
	}

	m.Unlock()

	// 前面对临界区的写操作, 需要加锁
	// 等待这个管道被其他地方写入
	ctx, cancel := context.WithTimeout(context.Background(), maxWaitDuration)
	defer cancel()
	select {
	case <-ctx.Done():
		log.Println("time out...")
		return -1, ctx.Err()
	case <-ch:
		// 啥也不用做, 进入临界区即可
	}

	m.Lock() // 避免二次修改了
	v = m.mp[k]
	m.Unlock()

	return v, nil
}

func RunRWMap() {
	m := NewRWMap()

	wg := sync.WaitGroup{}

	for i := 0; i < 4; i++ {
		go func(index int) {
			m.Put(index, index+4)
			wg.Add(1)
		}(i)
	}

	for i := 0; i < 5; i++ {
		go func(index int) {
			res, err := m.Get(index, 200*time.Second)
			if err != nil {
				panic(err)
			} else {
				log.Println("res: ", res)
				wg.Done()
			}
		}(i)
	}

	time.Sleep(2 * time.Second)
	m.Put(4, 4+4)
	wg.Add(1)

	wg.Wait()

	log.Println("all finished!")

}
