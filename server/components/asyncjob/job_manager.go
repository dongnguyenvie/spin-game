package asyncjob

import (
	"context"
	"fmt"
	"sync"
)

type group struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	return &group{
		isConcurrent: isConcurrent,
		jobs:         jobs,
		wg:           new(sync.WaitGroup),
	}
}

func (g *group) Run(ctx context.Context) error {

	errChan := make(chan error, len(g.jobs))

	for _, job := range g.jobs {
		if g.isConcurrent {
			go func(aj Job) {
				defer func() {
					if err := recover(); err != nil {
						fmt.Println("Error: ", err)
					}
				}()
				errChan <- g.runJob(ctx, aj)
			}(job)
		} else {
			err := g.runJob(ctx, job)

			if err != nil {
				return err
			}

			errChan <- err
		}
	}

	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			return v
		}
	}

	return nil
}

func (g *group) runJob(ctx context.Context, j Job) error {
	err := j.Execute(ctx)
	if err != nil {
		for {
			fmt.Println("Error", err)
			if j.State() == StateFailed {
				return err
			}

			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}

	return nil
}
