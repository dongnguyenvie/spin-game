package asyncjob

import (
	"context"
	"testing"
)

func TestSingleJob(t *testing.T) {
	var job1 Job = NewJob(nil, func(ctx context.Context) error {
		// do something
		return nil
	})

	err := job1.Execute(context.Background())
	if err != nil {
		t.Error("Execute failed")
	}

	if JobState.String(job1.State()) != "Completed" {
		t.Error("job was not completed yet")
	}

}
