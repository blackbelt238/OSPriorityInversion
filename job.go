package main

import (
	"fmt"
)

// Kind is the job type
type kind int

const (
	t1 kind = iota
	t2
	t3
)

// job represents a job to be processed
type job struct {
	index int  // index of the job in the job queue
	tarr  int  // time this job arrived at the queue
	trem  int  // the remaining time until the job is fully processed
	pri   int  // priority of the job
	kind  kind // the kind of the job
}

// createJob makes a new job
func createJob(i int, arr int, pri int, kind kind) *job {
	rem := 3 // initialize the remaining time to 3ms (for T1 and T3)
	// if a T2 is being created, adjust the remaining time accordingly
	if kind == t2 {
		rem = 10
	}
	return &job{i, arr, rem, pri, kind}
}

// Preempts evaluates whether the job preempts a given job
func (j *job) preempts(job *job) bool {
	if j.pri > job.pri && !(j.kind == t1 && job.kind == t3) {
		return true
	}
	return false
}

// Priority returns the priority of the job
func (j *job) priority() int {
	return j.pri
}

// String returns the string version of the job
func (j *job) string() string {
	return fmt.Sprintf("{%d %d %d %s}", j.tarr, j.trem, j.pri, kindToString(j.kind))
}

func kindToString(k kind) string {
	switch k {
	case 0:
		return "T1"
	case 1:
		return "T2"
	case 2:
		return "T3"
	}
	return "ERROR"
}
