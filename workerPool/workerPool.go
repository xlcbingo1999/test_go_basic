package workerPool

type WorkerPool struct {
	Size        int
	JobQueue    chan Job
	WorkerQueue chan *Worker // 用一个管道表示woeker队列
}

func NewWorkerPool(poolSize int, jobQueueLen int) *WorkerPool {
	return &WorkerPool{
		Size:        poolSize,
		JobQueue:    make(chan Job, jobQueueLen), // 建立了一个具有缓冲区的管道
		WorkerQueue: make(chan *Worker, poolSize),
	}
}

func (wp *WorkerPool) GetJobQueue() <-chan Job {
	return wp.JobQueue
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.Size; i++ {
		worker := NewWorker()
		worker.Start(wp)
	}

	go func() {
		for job := range wp.GetJobQueue() { // 这是Go编译器更推荐的写法
			worker := <-wp.WorkerQueue // 从worker池中拿到一个worker
			worker.JobQueue <- job     // 把job写到worker的JobQueue这个管道中, 模拟的就是队列
		}
	}()

}
