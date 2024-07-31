package mutexstruct

import "sync"

// RWMutexMap 是一个简单的 map + sync.RWMutex 的并发安全散列表实现
type RWMutexMap struct {
	data map[interface{}]interface{}
	mu   sync.RWMutex
}

// 全部都是将接收者定义为指针的
func (m *RWMutexMap) Load(k interface{}) (v interface{}, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok = m.data[k]
	return
}

func (m *RWMutexMap) Store(k, v interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}
