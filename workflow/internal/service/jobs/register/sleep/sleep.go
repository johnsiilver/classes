/*
Package sleep registers a job that is used to sleep between Jobs.

Register name: "sleep"
Args:
	"seconds"(mandatory): The amount of time to sleep in seconds. Must be >= 1.
Result:
	Sleeps for the amount of time specified.
*/
package sleep

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/johnsiilver/classes/workflow/internal/service/jobs"

	pb "github.com/johnsiilver/classes/workflow/proto"
)

// This registers our Job on server startup.
func init() {
	jobs.Register("sleep", newJob())
}

type args struct {
	d time.Duration
}

func (a *args) validate(args map[string]string) error {
	must := map[string]bool{
		"seconds": false,
	}

	for k, v := range args {
		switch k {
		case "seconds":
			i, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("arg(seconds) is not an integer(%s)", v)
			}
			if i < 1 {
				return fmt.Errorf("arg(seconds) cannot be less than 1(%d)", i)
			}
			must["seconds"] = true
			a.d = time.Duration(i) * time.Second
		default:
			return fmt.Errorf("validateDecom had invalid arg(%s)", k)
		}
	}

	for k, v := range must {
		if !v {
			return fmt.Errorf("missing required arg(%s)", k)
		}
	}
	return nil
}

// Job implements jobs.Job.
type Job struct {
	args args
}

func newJob() *Job {
	return &Job{}
}

// Validate implements jobs.Job.Validate().
func (j *Job) Validate(job *pb.Job) error {
	a := args{}
	if err := a.validate(job.Args); err != nil {
		return err
	}
	j.args = a
	return nil
}

// Run implements jobs.Job.Run().
func (j *Job) Run(ctx context.Context, job *pb.Job) error {
	time.Sleep(j.args.d)
	return nil
}
