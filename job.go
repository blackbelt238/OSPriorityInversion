package main

import (
	"fmt"
)

// Job represents a job to be processed
type Job struct {
	tarr  int // time this job arrived at the queue
	tproc int // time this job requires to be fully processed
	pri   int // priority of the job
}

// CreateFillerJob returns a new, invalid job
func CreateFillerJob() *Job {
	return &Job{-1, -1, -1}
}

// IsFiller checks to see if the given job is invalid
func (j *Job) IsFiller() bool {
	return j.tarr == -1 && j.tproc == -1 && j.pri == -1
}

// Priority returns the priority of the job
func (j *Job) Priority() int {
	return j.pri
}

// String returns the string version of the job
func (j *Job) String() string {
	return fmt.Sprintf("%d %d %d", j.tarr, j.tproc, j.pri)
}
