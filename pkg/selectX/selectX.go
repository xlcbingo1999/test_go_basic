package selectX

import (
	"fmt"
	"strconv"
	"sync"
)

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select { // 用select代替switch
// 		case c <- x: // 向管道写入x
// 			x, y = y, x+y
// 		case <-quit: // 读取管道的quit
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

type Task struct {
	val int
}

type errorString struct {
	s string
}

func (e *errorString) Error() string { // 只要实现了Error()就可以重写
	return e.s
}

func (t *Task) Run() error {
	if t.val > 5 {
		return &errorString{strconv.Itoa(t.val)}
	}
	fmt.Println("result: ", t.val)

	return nil
}

func runAllTasks() {
	tasks := make([]*Task, 0)
	for i := 0; i < 7; i++ {
		tasks = append(tasks, &Task{val: i})
	}
	errCh := make(chan error, len(tasks))

	wg := sync.WaitGroup{} // 用来等待一批Goroutine结束
	wg.Add(len(tasks))     // 使用Add()添加需要等待的个数

	for i := range tasks {
		go func(index int) {
			defer wg.Done() // 使用Done()来表示完成一个Goroutine
			if err := tasks[index].Run(); err != nil {
				errCh <- err
			}
		}(i)
	}

	wg.Wait() // 等待同步完成, 这里只申请了一个WaitGroup, 所以就一个Wait()即可, 会一直等到完成的

	select {
	case err := <-errCh:
		fmt.Println("err: ", err)
	default:
		fmt.Println("all success")
	}
}

func nonWaitSelectCase() {
	ch := make(chan int)

	select {
	case i := <-ch: // 从管道中获取数据, 并写到i中
		fmt.Println(i)
	default: // 非阻塞, 如果此时管道还没有值, 就应该让协程的其他流程继续执行, 不应该强行占着
		fmt.Println("default")
	}
}

func RunSelect() {
	// c := make(chan int)
	// quit := make(chan int)
	// fibonacci(c, quit)

	// time.Sleep(100 * time.Millisecond)
	// quit <- 1

	// fmt.Printf("\n\n\n\n")

	nonWaitSelectCase()

	runAllTasks()
}
