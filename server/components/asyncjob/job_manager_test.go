package asyncjob

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestGroupJob(t *testing.T) {
	var job1 Job = NewJob(nil, func(ctx context.Context) error {
		fmt.Println("job 1 done")
		return nil
	})

	var job2 Job = NewJob(nil, func(ctx context.Context) error {
		fmt.Println("job 2 done")
		return nil
	})

	var job3 Job = NewJob(nil, func(ctx context.Context) error {
		fmt.Println("job 3 done")
		return nil
	})

	group := NewGroup(true, job1, job2, job3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
