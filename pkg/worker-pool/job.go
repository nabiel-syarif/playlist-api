package workerpool

import "context"

type ExecFunc func(ctx context.Context, args interface{}) (interface{}, error)

type Job struct {
	Id   int
	Fn   ExecFunc
	Args interface{}
}

type JobResult struct {
	JobId int
	Value interface{}
	Error error
}

func (job *Job) execute(ctx context.Context) JobResult {
	value, err := job.Fn(ctx, job.Args)
	if err != nil {
		return JobResult{
			JobId: job.Id,
			Error: err,
		}
	}
	return JobResult{
		JobId: job.Id,
		Value: value,
	}
}
