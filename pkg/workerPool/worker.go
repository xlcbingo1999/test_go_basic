package workerPool

type Worker struct {
	JobQueue chan Job  // 用一个管道去表示任务缓冲队列
	Quit     chan bool // 表示Worker的退出的状态量
}

func NewWorker() Worker {
	return Worker{
		JobQueue: make(chan Job),
		Quit:     make(chan bool),
	}
}

func (w *Worker) Start(workerPool *WorkerPool) {
	go func() {
		for {
			workerPool.WorkerQueue <- w // 把自己写入到worker队列中
			select {
			case job := <-w.JobQueue:
				job.Run()
			case <-w.Quit: // worker的结束流程
				return
			}
		}
	}()
}
