package asyncjob

import (
	"context"
	"time"
)

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDurations(times []time.Duration)
}

type JobState int

func (state JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[state]
}

type JobHandler func(ctx context.Context) error

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

type JobConfig interface{}

type job struct {
	config     JobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(config JobConfig, handler JobHandler) *job {
	return &job{config: config, handler: handler, retryIndex: -1, state: StateInit, stopChan: make(chan bool)}
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateInit

	err := j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	return nil
}

func (j *job) State() JobState {
	return j.state
}

func (j *job) SetRetryDurations(times []time.Duration) {
}
