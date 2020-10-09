package main

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type worker struct {
	workerPool chan *worker
	jobChannel chan Job
	stop       chan struct{}
}

func (w *worker) start() {
	go func() {
		var job Job
		for {
			// worker free, add it to pool
			select {
			case job = <-w.jobChannel:
				job()
			case <-w.stop:
				w.stop <- struct{}{}
				return
			}
			w.workerPool <- w

		}
	}()
}

func newWorker(pool chan *worker) *worker {
	return &worker{
		workerPool: pool,
		jobChannel: make(chan Job),
		stop:       make(chan struct{}),
	}
}

type dispatcher struct {
	workerPool chan *worker
	jobQueue   chan Job
	stop       chan struct{}
	pool       *Pool
	once       sync.Once
}

func (d *dispatcher) dispatch() {
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case job := <-d.jobQueue:
			select {
			case worker := <-d.workerPool:
				worker.jobChannel <- job
			default:
				runningWorker := atomic.LoadInt64(&d.pool.runningWorker)
				if runningWorker < d.pool.maxWorkerSize {
					atomic.AddInt64(&d.pool.runningWorker, 1)
					newWorker := newWorker(d.workerPool)
					newWorker.start()
					newWorker.jobChannel <- job

				} else {
					select {
					case worker := <-d.workerPool:
						worker.jobChannel <- job
					case <-d.stop:
						for i := 0; i < len(d.workerPool); i++ {
							worker := <-d.workerPool
							worker.stop <- struct{}{}
							<-worker.stop
						}
						d.stop <- struct{}{}
						return

					}
				}
			}

		case <-timer.C:
			if len(d.jobQueue) < len(d.workerPool) {
				select {
				case w := <-d.workerPool:
					atomic.AddInt64(&d.pool.runningWorker, -1)
					w.stop <- struct{}{}
					<-w.stop
					w.jobChannel = nil
					w.stop = nil
					w = nil

				default:

				}
			}
		case <-d.stop:
			d.once.Do(func() {
				for {
					runningWork := atomic.LoadInt64(&d.pool.runningWorker)
					if len(d.workerPool) > 0 || runningWork > 0 {
						for i := 0; i < len(d.workerPool); i++ {
							worker := <-d.workerPool
							worker.stop <- struct{}{}
							<-worker.stop
							atomic.AddInt64(&d.pool.runningWorker, -1)
						}
					} else {
						return
					}
				}
			})
			d.stop <- struct{}{}
			return
		}
	}
}

func newDispatcher(workerPool chan *worker, jobQueue chan Job, pool *Pool) *dispatcher {
	d := &dispatcher{
		workerPool: workerPool,
		jobQueue:   jobQueue,
		stop:       make(chan struct{}),
		pool:       pool,
	}
	go d.dispatch()
	return d
}
func (pool *Pool) Running() int64 {
	runningWorker := atomic.LoadInt64(&pool.runningWorker)
	return runningWorker - int64(len(pool.dispatcher.workerPool))
}

type Job func()

type Pool struct {
	JobQueue      chan Job
	dispatcher    *dispatcher
	wg            sync.WaitGroup
	runningWorker int64
	maxWorkerSize int64
	maxJobSize    int64
}

func NewPool(numWorkers int, jobQueueLen int) *Pool {
	maxWorkerSize := 100000
	maxJobSize := 100000
	if numWorkers < 0 {
		panic(errors.New("numWorkers is negative"))
	}
	if numWorkers < maxWorkerSize {
		maxWorkerSize = numWorkers
	}
	if jobQueueLen < 0 {
		panic(errors.New("jobQueueLen is negative"))
	}
	if jobQueueLen < maxJobSize {
		maxJobSize = jobQueueLen
	}

	jobQueue := make(chan Job, maxJobSize)
	workerPool := make(chan *worker, maxWorkerSize)
	pool := &Pool{
		JobQueue:      jobQueue,
		maxWorkerSize: int64(maxWorkerSize),
		maxJobSize:    int64(maxJobSize),
	}
	pool.dispatcher = newDispatcher(workerPool, jobQueue, pool)
	return pool
}

func (p *Pool) JobDone() {
	p.wg.Done()
}

func (p *Pool) WaitCount(count int) {
	p.wg.Add(count)
}

func (p *Pool) WaitAll() {
	p.wg.Wait()
}

func (p *Pool) Release() {
	p.dispatcher.stop <- struct{}{}
	<-p.dispatcher.stop
}
