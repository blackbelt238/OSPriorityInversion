package main

import (
	"container/heap"
	"fmt"
)

type jobQueue []*job // a priority queue of jobs

type processor struct {
	que jobQueue // the queue of jobs in priority order
	bl  []*job   // the backlog of jobs
}

// createProcessor makes and returns a new processor out of new processor components
func createProcessor() *processor {
	que := make(jobQueue, 0, 5)
	heap.Init(&que)

	return &processor{que, make([]*job, 0)}
}

// AddJob takes in information for a job, creates it, and adds it to the processor
func (p *processor) addJob(tarr int, pri int) {
	var k kind
	switch pri {
	case 1:
		k = t3
	case 2:
		k = t2
	case 3:
		k = t1
	default:
		fmt.Println("ERROR: invalid kind")
	}
	p.bl = append(p.bl, createJob(len(p.bl), tarr, pri, k))
}

// run simulates a run on the processor based on what's been added to it
func (p *processor) run() {
	p.printBacklog()
	time := 0 // time of processor run (in ms)
}

// PrintBacklog prints the job backlog of the processor
func (p *processor) printBacklog() {
	fmt.Print("Job Backlog: ")
	for _, job := range p.bl {
		fmt.Print(job.string() + " ")
	}
	fmt.Println()
}

// Len implements sort.Len
func (jq jobQueue) Len() int { return len(jq) }

// Less implements sort.Less
func (jq jobQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return jq[i].pri > jq[j].pri
}

// Swap implements sort.Swap
func (jq jobQueue) Swap(i, j int) {
	jq[i], jq[j] = jq[j], jq[i]
	jq[i].index = i
	jq[j].index = j
}

// Pop implements heap.Pop
func (jq *jobQueue) Pop() interface{} {
	old := *jq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*jq = old[0 : n-1]
	return item
}

// Push implements heap.Push
func (jq *jobQueue) Push(x interface{}) {
	n := len(*jq)
	job := x.(*job)
	job.index = n
	*jq = append(*jq, job)
}
