package workerPool2

import (
	"fmt"
	"time"
)

// type Job struct {
// 	Index     int
// 	WaitGroup *sync.WaitGroup
// }

// func (j *Job) RunJob() {
// 	fmt.Println("run job ", j.Index)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("finished job ", j.Index)
// 	j.WaitGroup.Done()
// }

// // worker一般维护一个本地的任务队列, 然后只要新建了这个worker就要不断for-select执行本地任务队列里面的任务
// // 当一个任务被分配给worker的时候，就会消费掉workerPool的一个worker, 当任务完成的时候就要把这个worker重新生产回去
// type Worker struct {
// 	Index         int
// 	LocalJobQueue chan *Job // Woker需要设计一个本地的执行Job队列, 暂时不需要设计缓冲区
// 	Quit          chan bool
// }

// func GetWorker(index int) *Worker {
// 	return &Worker{
// 		Index:         index,
// 		LocalJobQueue: make(chan *Job),
// 		Quit:          make(chan bool),
// 	}
// }

// func (w *Worker) RunLocalJob(wp *WorkerPool) {
// 	go func() {
// 		for {
// 			select {
// 			case job := <-w.LocalJobQueue:
// 				job.RunJob()
// 				// 当完成任务之后, 需要重新把这个worker生产回来
// 				wp.WorkerQueue <- w

// 			case <-w.Quit:
// 				fmt.Println("Quit Worker ", w.Index)
// 				return
// 			}
// 		}
// 	}()
// }

// type WorkerPool struct {
// 	Size        int
// 	JobQueue    chan *Job
// 	WorkerQueue chan *Worker
// }

// func GetWorkerPool(maxJobNum int, maxWorkerNum int) *WorkerPool {
// 	pool := &WorkerPool{
// 		Size:        maxWorkerNum,
// 		JobQueue:    make(chan *Job, maxJobNum),
// 		WorkerQueue: make(chan *Worker, maxWorkerNum),
// 	}
// 	for i := 0; i < maxWorkerNum; i++ {
// 		pool.AddWorker(i)
// 	}
// 	return pool
// }

// func (wp *WorkerPool) AddWorker(index int) {
// 	worker := GetWorker(index)
// 	worker.RunLocalJob(wp)
// 	wp.WorkerQueue <- worker // 只生产了两次
// }

// func (wp *WorkerPool) Run() {
// 	// 不断循环任务等待队列中的job, 然后找到对应的执行者
// 	go func() {
// 		for waitJob := range wp.JobQueue {
// 			// 读取job
// 			// fmt.Println("get waitJob: ", waitJob.Index)
// 			targetWorker := <-wp.WorkerQueue // 没有内容可以消费
// 			targetWorker.LocalJobQueue <- waitJob
// 		}
// 	}()
// }

func RunWorkerPool() {
	ch := make(chan int)
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("wait for pool")
	case i := <-ch:
		fmt.Println("i: ", i)
	}
	fmt.Println("out")

	ch <- 2
	// maxWorkerNum := 2
	// maxJobNum := 10
	// waitgroup := &sync.WaitGroup{}
	// waitgroup.Add(maxJobNum)

	// workerPool := GetWorkerPool(
	// 	maxJobNum, maxWorkerNum,
	// )
	// workerPool.Run()

	// // 发送任务
	// go func() {
	// 	for i := 0; i < maxJobNum; i++ {
	// 		workerPool.JobQueue <- &Job{
	// 			Index:     i,
	// 			WaitGroup: waitgroup,
	// 		}
	// 	}
	// }()

	// waitgroup.Wait()

	// defer func() {
	// 	for {
	// 		select { // select外部一般要加一个无限的for循环
	// 		case worker := <-workerPool.WorkerQueue:
	// 			worker.Quit <- true
	// 		default:
	// 			fmt.Println("All finished!")
	// 			return
	// 		}
	// 	}
	// }()
}
