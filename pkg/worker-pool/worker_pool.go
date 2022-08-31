package workerpool

import (
	"context"
	"log"
	"sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, jobsChan <-chan Job, resultsChan chan<- JobResult) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobsChan:
			if !ok {
				return
			}
			resultsChan <- job.execute(ctx)
		case <-ctx.Done():
			log.Printf("worker cancelled because context done. Err : %v\n", ctx.Err())
			return
		}
	}
}

type WorkerPool struct {
	WorkerCounts int
	resultsChan  chan JobResult
	jobsChan     chan Job
	wg           sync.WaitGroup
}

func New(numOfWorkers int) *WorkerPool {
	return &WorkerPool{
		resultsChan:  make(chan JobResult, numOfWorkers),
		jobsChan:     make(chan Job, numOfWorkers),
		WorkerCounts: numOfWorkers,
	}
}

func (w *WorkerPool) AddJob(job Job) {
	w.jobsChan <- job
}

func (w *WorkerPool) FromJobs(jobs []Job) {
	for _, job := range jobs {
		w.jobsChan <- job
	}
}

func (w *WorkerPool) Close() {
	close(w.jobsChan)
}

func (w *WorkerPool) Run(ctx context.Context) {
	log.Println("WorkerPool running")
	for i := 0; i < w.WorkerCounts; i++ {
		w.wg.Add(1)
		go worker(ctx, &w.wg, w.jobsChan, w.resultsChan)
		log.Printf("Worker %d started\n", i+1)
	}
}

func (w *WorkerPool) IsWorkerPoolWaitingQueueEmpty() bool {
	return len(w.jobsChan) == 0
}

func (w *WorkerPool) Wait() {
	w.wg.Wait()
}

func (w *WorkerPool) Results() <-chan JobResult {
	return w.resultsChan
}
