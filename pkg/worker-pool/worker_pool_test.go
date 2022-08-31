package workerpool

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	workerCount := 3
	workerPool := New(workerCount)
	require.NotEmpty(t, workerPool)
	require.Equal(t, workerCount, workerPool.WorkerCounts)
	require.Equal(t, workerCount, cap(workerPool.jobsChan))
	require.Equal(t, workerCount, cap(workerPool.resultsChan))
}

func TestWorkerPool_AddJob(t *testing.T) {
	workerCount := 1
	workerPool := New(workerCount)
	workerPool.AddJob(Job{
		Id: 1,
		Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
			log.Println("Job executed")
			return true, nil
		},
	})

	require.Equal(t, 1, len(workerPool.jobsChan))
	job := <-workerPool.jobsChan
	require.NotEmpty(t, job)
	require.Equal(t, 1, job.Id)
}

func TestWorkerPool_FromJobs(t *testing.T) {
	jobsCount := 3
	workerPool := New(jobsCount)

	jobs := make([]Job, jobsCount)
	for i := 0; i < jobsCount; i++ {
		jobs[i] = Job{
			Id: i,
			Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
				return i, nil
			},
			Args: nil,
		}
	}

	workerPool.FromJobs(jobs)

	require.Equal(t, jobsCount, len(workerPool.jobsChan))
	for i := 0; i < jobsCount; i++ {
		job := <-workerPool.jobsChan
		require.Equal(t, i, job.Id)
	}
}

func TestWorkerPool_Close(t *testing.T) {
	workerPool := New(1)
	workerPool.AddJob(Job{
		Id: 1,
		Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
			return 1, nil
		},
	})
	workerPool.Close()
	for i := 0; i < 2; i++ {
		job, isOpen := <-workerPool.jobsChan
		if i == 0 {
			require.Equal(t, 1, job.Id)
			require.True(t, isOpen)
		} else {
			require.Equal(t, 0, job.Id)
			require.False(t, isOpen)
		}
	}
}

func TestWorkerPool_Run(t *testing.T) {
	workerCount := 1
	workerPool := New(workerCount)
	jobsCount := 5
	jobs := make([]Job, jobsCount)
	for i := 0; i < jobsCount; i++ {
		jobs[i] = Job{
			Id: i,
			Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
				return i, nil
			},
		}
	}

	go func() {
		workerPool.FromJobs(jobs)
		workerPool.Close()
	}()
	
	workerPool.Run(context.Background())
	results := workerPool.Results()

	for i := 0; i < jobsCount; i++ {
		job := <-results
		require.NotEmpty(t, job)
	}
	workerPool.Wait()
}

func TestWorker_ContextError(t *testing.T) {
	workerPool := New(2)
	job := Job{
		Id: 1,
		Fn: func(ctx context.Context, args interface{}) (interface{}, error) {
			return 1, nil
		},
	}

	durationTimeout := time.Millisecond * 500
	ctx, cancel := context.WithTimeout(context.Background(), durationTimeout)
	defer cancel()
	results := workerPool.Results()

	go func() {
		<-time.After(time.Second * 1)
		// should never be called because context already timeout
		workerPool.AddJob(job)
		workerPool.Close()
	}()
	// in Run(), go routine waiting for context timeout or job sent through the channel via AddJob() or FromJobs() function
	workerPool.Run(ctx)

	<-time.After(durationTimeout * 2)
	require.Len(t, results, 0)
}

func TestWorkerPool_Results(t *testing.T) {
	workerCount := 1
	workerPool := New(workerCount)
	results := workerPool.Results()

	require.IsType(t, make(<-chan JobResult, workerCount), results)
	require.Equal(t, workerCount, cap(results))
}
