package lfu

import (
	"fmt"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

type LFUCache struct {
	freq2keys map[int]linkedhashmap.Map
	key2val   map[int]int
	key2freq  map[int]int
	minFreq   int
	cap       int
	count     int
}

func NewLFUCache(cap int) *LFUCache {
	cache := &LFUCache{
		freq2keys: make(map[int]linkedhashmap.Map),
		key2val:   make(map[int]int),
		key2freq:  make(map[int]int),
		minFreq:   0,
		cap:       cap,
		count:     0,
	}
	return cache
}

func (c *LFUCache) updateExistKeyFreq(key int) {
	// 更新
	originFreq := c.key2freq[key]
	newFreq := originFreq + 1
	c.key2freq[key] = newFreq
	originMap := c.freq2keys[originFreq]
	originMap.Remove(key)

	newMap, exist := c.freq2keys[newFreq]
	if !exist {
		newMap = *linkedhashmap.New()
		newMap.Put(key, 0)
		c.freq2keys[newFreq] = newMap
	} else {
		newMap.Put(key, 0)
	}

	if originMap.Empty() { // 这个代码就很大的问题, 很烦不想改了
		delete(c.freq2keys, originFreq)
		if originFreq == c.minFreq {
			c.minFreq = newFreq
		}
	}
}

func (c *LFUCache) get(key int) int {
	val, exist := c.key2val[key]
	if !exist {
		return -1
	}

	c.updateExistKeyFreq(key)
	return val
}

func (c *LFUCache) put(key int, val int) {
	_, exist := c.key2val[key]
	if !exist {
		// 当有一个新的进来且超过容量的时候
		// 1. 需要先剔除掉最近最少使用的元素
		// 2. 把新进来的元素设置为频率1, 然后增加进去
		if c.count >= c.cap {
			// 删除即可
			smallMap := c.freq2keys[c.minFreq]
			mapIter := smallMap.Iterator()

			if mapIter.Next() { // 迭代器走到了下一个格子, 直接就可以处理了
				deleteKey := mapIter.Key()
				smallMap.Remove(deleteKey)
				delete(c.key2val, deleteKey.(int))
				delete(c.key2val, deleteKey.(int))
				c.count -= 1
			}
		}
		c.key2val[key] = val
		c.key2freq[key] = 1
		c.minFreq = 1
		c.count += 1
	} else {
		c.key2val[key] = val
		c.updateExistKeyFreq(key)
	}
}

func (c *LFUCache) Print() {
	for k, v := range c.key2val {
		fmt.Printf("key: %d => value: %d\n", k, v)
	}
	fmt.Println("\n\n\n")
}

func RunLFU() {
	cache := NewLFUCache(2)
	cache.put(1, 2)
	cache.Print()
	cache.put(3, 4)
	cache.Print()
	fmt.Println(cache.get(1))
	cache.Print()
	fmt.Println(cache.get(3))
	cache.Print()
	fmt.Println(cache.get(2))
	cache.Print()
	cache.put(3, 4)
	cache.Print()
	cache.put(2, 5)
	cache.Print()
	fmt.Println(cache.get(2))
	cache.Print()
}
