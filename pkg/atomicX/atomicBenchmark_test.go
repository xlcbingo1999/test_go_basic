package atomicX

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Config struct {
	sync.RWMutex // 这是一个读写互斥锁实例
	endpoint     string
}

func BenchmarkPMutexSet(b *testing.B) {
	config := Config{}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Lock()
			config.endpoint = "api.example.com"
			config.Unlock()
		}
	})
}

func BenchmarkPMutexGet(b *testing.B) {
	config := Config{endpoint: "api.example.com"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.RLock()
			_ = config.endpoint
			config.RUnlock()
		}
	})
}

func BenchmarkPAtomicSet(b *testing.B) {
	var config atomic.Value
	c := Config{endpoint: "api.example.com"}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			config.Store(c.endpoint)
		}
	})
}

func BenchmarkPAtomicGet(b *testing.B) {
	var config atomic.Value
	config.Store(Config{endpoint: "api.example.com"})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = config.Load()
		}
	})
}
