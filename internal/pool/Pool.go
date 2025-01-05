package pool

import (
	"errors"
	"log"
	"sync"
)

type Job func() error

type Pool struct {
	workQueue chan Job
	wg        sync.WaitGroup
}

// NewPool initializes a worker pool with a fixed number of workers
func NewPool(workerCount int) *Pool {
	pool := &Pool{
		workQueue: make(chan Job, workerCount*2),
	}

	// Start workers
	for i := 0; i < workerCount; i++ {
		go pool.worker(i)
	}

	return pool
}

// worker processes jobs from the queue
func (p *Pool) worker(id int) {
	for job := range p.workQueue {
		if err := job(); err != nil {
			log.Printf("Worker %d: Job failed with error: %v\n", id, err)
		}
	}
}

// AddJob submits a job to the pool
func (p *Pool) AddJob(job Job) error {
	select {
	case p.workQueue <- job:
		return nil
	default:
		return errors.New("worker pool is full, unable to accept new jobs")
	}
}

// Shutdown gracefully shuts down the worker pool
func (p *Pool) Shutdown() {
	close(p.workQueue)
	p.wg.Wait()
}
