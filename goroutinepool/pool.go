package goroutinepool

import (
	"errors"
	"sync"
	"sync/atomic"
)

// WorkerPool 协程池接口
type WorkerPool interface {
	// Run 启动协程池
	Run()
	// AddTask 添加任务到协程池
	AddTask(task func()) error
	// Shutdown 关闭协程池
	Shutdown()
}

// worker 协程池中的工作协程
type worker struct {
	workerPool chan chan func()
	taskChan   chan func()
	quit       chan struct{}
	wg         *sync.WaitGroup // 添加WaitGroup指针，用于任务完成时的通知
}

// newWorker 创建一个新的工作协程
func newWorker(workerPool chan chan func(), wg *sync.WaitGroup) *worker {
	return &worker{
		workerPool: workerPool,
		taskChan:   make(chan func()),
		quit:       make(chan struct{}),
		wg:         wg,
	}
}

// start 启动工作协程
func (w *worker) start() {
	go func() {
		for {
			// 将当前工作协程的任务通道注册到协程池
			w.workerPool <- w.taskChan
			select {
			case task, ok := <-w.taskChan:
				if !ok {
					return
				}
				// 执行任务
				task()
				w.wg.Done() // 任务执行完毕后，通知WaitGroup
			case <-w.quit:
				// 退出工作协程
				return
			}
		}
	}()
}

// stop 停止工作协程
func (w *worker) stop() {
	close(w.quit)
}

// EnhancedWorkerPool 增强的协程池实现
type EnhancedWorkerPool struct {
	workerPool chan chan func()
	workers    []*worker
	taskQueue  chan func()
	quit       chan struct{}
	shutdown   atomic.Bool
	wg         sync.WaitGroup // 用于跟踪未完成的任务
}

// NewWorkerPool 创建一个新的增强协程池
func NewWorkerPool(numWorkers, taskQueueLen int) *EnhancedWorkerPool {
	pool := &EnhancedWorkerPool{
		workerPool: make(chan chan func(), numWorkers),
		workers:    make([]*worker, numWorkers),
		taskQueue:  make(chan func(), taskQueueLen),
		quit:       make(chan struct{}),
	}

	// 启动工作协程
	for i := 0; i < numWorkers; i++ {
		worker := newWorker(pool.workerPool, &pool.wg)
		pool.workers[i] = worker
		worker.start()
	}

	return pool
}

// Run 启动协程池
func (p *EnhancedWorkerPool) Run() {
	go func() {
		for {
			select {
			case task := <-p.taskQueue:
				// 从工作协程池中获取一个空闲的工作协程
				workerTaskChan := <-p.workerPool
				// 将任务发送到工作协程
				workerTaskChan <- task
			case <-p.quit:
				// 关闭所有工作协程
				for _, worker := range p.workers {
					worker.stop()
				}
				return
			}
		}
	}()
}

// AddTask 添加任务到协程池
func (p *EnhancedWorkerPool) AddTask(task func()) error {
	if p.shutdown.Load() {
		return errors.New("worker pool is shutdown")
	}

	p.wg.Add(1) // 增加等待组的计数
	select {
	case p.taskQueue <- task:
		// 任务添加到队列成功
	case <-p.quit:
		// 如果协程池已经关闭，则返回错误
		p.wg.Done() // 减少等待组的计数
		return errors.New("worker pool is shutdown")
	}
	return nil
}

// Shutdown 关闭协程池
func (p *EnhancedWorkerPool) Shutdown() {
	p.shutdown.Store(true) // 设置shutdown标志
	p.wg.Wait()            // 等待所有任务处理完毕
	close(p.quit)          // 关闭quit通道，通知所有工作协程停止
}
