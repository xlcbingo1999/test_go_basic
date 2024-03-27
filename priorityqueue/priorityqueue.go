package priorityqueue

import (
	"container/heap"
	"log"
)

type Item struct {
	val      int
	priority int
	index    int
}

type PriorityQueue []*Item // 核心的是[], 只是它实现了接口, 完成了多态

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority // 这里取反就是大根堆
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *PriorityQueue) Push(nitem interface{}) {
	n := len(*pq)
	item := nitem.(*Item)
	item.index = n
	*pq = append(*pq, item)
	pq.Update(item, item.val, item.priority)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1] // 切片操作进行赋值，只会处理最后一个元素
	item.index = -1
	return item
}

func (pq *PriorityQueue) Update(nitem interface{}, val int, priority int) {
	item := nitem.(*Item)
	item.val = val
	item.priority = priority
	heap.Fix(pq, item.index) // 传进去用于修改index对应的位置元素的情况
}

func RunPriorityQueue() {
	items := map[int]int{2: 3, 3: 2, 5: 1, 8: 4}
	pq := make(PriorityQueue, len(items))
	ind := 0
	for k, v := range items {
		pq[ind] = &Item{
			val:      k,
			priority: v,
			index:    ind,
		}
		ind += 1
	}

	heap.Init(&pq) // 初始化堆, 本身就带着排序的功能
	item := &Item{
		val:      100,
		priority: 10000,
		index:    ind,
	}
	heap.Push(&pq, item)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		log.Println("item: ", item.priority, " => ", item.val, " => index: ", item.index)
	}
}
