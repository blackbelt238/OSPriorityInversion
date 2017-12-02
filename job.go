package main

import (
	"fmt"
)

// Kind is the job type
type Kind int

const (
	T1 Kind = iota
	T2
	T3
)

// Job represents a job to be processed
type Job struct {
	tarr int  // time this job arrived at the queue
	trem int  // the remaining time until the job is fully processed
	pri  int  // priority of the job
	kind Kind // the kind of the job
}

// CreateJob makes a new job
func CreateJob(arr int, pri int, kind Kind) *Job {
	rem := 3 // initialize the remaining time to 3ms (for T1 and T3)
	// if a T2 is being created, adjust the remaining time accordingly
	if kind == T2 {
		rem = 10
	}
	return &Job{arr, rem, pri, kind}
}

// Preempts evaluates whether the job preempts a given job
func (j *Job) Preempts(job *Job) bool {
	if j.pri > job.pri && !(j.kind == T1 && job.kind == T3) {
		return true
	}
	return false
}

// Priority returns the priority of the job
func (j *Job) Priority() int {
	return j.pri
}

// String returns the string version of the job
func (j *Job) String() string {
	return fmt.Sprintf("%d %d %d", j.tarr, j.trem, j.pri)
}
