package main

import (
	"fmt"
)

// Kind is the job type
type priority int

const (
	t1 priority = 3
	t2 priority = 2
	t3 priority = 1
)

// job represents a job to be processed
type job struct {
	index int      // index of the job in the job queue
	tarr  int      // time this job arrived at the queue
	trem  int      // the remaining time until the job is fully processed
	pri   priority // priority of the job
}

// createJob makes a new job
func createJob(i int, arr int, pri priority) *job {
	rem := 3 // initialize the remaining time to 3ms (for T1 and T3)
	// if a T2 is being created, adjust the remaining time accordingly
	if pri == t2 {
		rem = 10
	}
	return &job{i, arr, rem, pri}
}

func (j *job) kind() string {
	switch j.pri {
	case t1:
		return "T1"
	case t2:
		return "T2"
	case t3:
		return "T3"
	}
	panic("ERROR: Job has invalid kind.")
}

func (j *job) printVal() string {
	switch j.pri {
	case t1:
		return "1"
	case t2:
		return "N"
	case t3:
		return "3"
	}
	panic("Invalid job priority")
}

// String returns the string version of the job
func (j *job) string() string {
	return fmt.Sprintf("{%d %d %s}", j.tarr, j.trem, j.kind())
}

func (pri priority) canPreempt(pri2 priority) bool {
	if int(pri) > int(pri2) {
		if pri == t1 && pri2 == t3 {
			return false
		}
		return true
	}
	return false
}
