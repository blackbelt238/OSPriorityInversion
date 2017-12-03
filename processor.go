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

// AddJob takes in information for a job, creates it, and adds it to the processor's backlog
func (p *processor) addJob(tarr int, pri priority) {
	p.bl = append(p.bl, createJob(len(p.bl), tarr, pri))
}

// run simulates a run on the processor based on what's been added to the backlog at runtime
func (p *processor) run() {
	p.printBacklog()
	time := 0 // time of processor run (in ms)

	// the current job's index and kind
	var cind int
	var ckind string

	// the previous job's index and kind
	pind := -1
	pkind := ""

	// continue processing as long as there are jobs to process
	for time < 50 && (p.que.Len() > 0 || len(p.bl) > 0) {
		// fmt.Println("time =", time)
		// if a job has arrived
		if len(p.bl) > 0 && p.bl[0].tarr == time {
			// if p.que.Len() == 0 || p.bl[0].pri.canPreempt(p.que[p.que.Len()-1].pri) {
			// fmt.Println("pulled from backlog to queue:", p.bl[0].kind())
			p.que.Push(p.bl[0]) // add the job to the priority queue
			p.bl = p.bl[1:]     // remove it from the backlog

			// }
			// fmt.Println("cannot preempt")
		}

		// if there is a job to work on
		if p.que.Len() > 0 {
			cind = p.que[p.que.Len()-1].index
			ckind = p.que[p.que.Len()-1].kind()

			// if preemption occurred
			if cind != pind {
				// if this is not the first iteration
				if pind != -1 {
					fmt.Println(" " + pkind + ".") // end processing on the previous process
				}
				// start the next process
				fmt.Print("time " + fmt.Sprintf("%v", time) + ", " + ckind + " ")
				pind = cind
				pkind = ckind
			}

			// make progress on the current job
			fmt.Print(p.que[p.que.Len()-1].printVal())
			// fmt.Println(p.que[p.que.Len()-1].string())
			p.que[p.que.Len()-1].trem--

			// if work has finished on the job, remove it from the processor
			// the next pass through will detect a change and place the ending characters properly
			if p.que[p.que.Len()-1].trem == 0 {
				p.que.Pop()
			}
		}

		time++ // progress the simulation
	}
	fmt.Println(" " + pkind + ".")
}

// PrintBacklog prints the job backlog of the processor
func (p *processor) printBacklog() {
	fmt.Print("inputs: ")
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
