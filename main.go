package main

// runs the simulation
func main() {
	// the 3 priority levels
	T1 := 3
	T2 := 2
	T3 := 1

	p := createProcessor()
	p.addJob(1, T3)
	p.addJob(2, T2)
	p.addJob(6, T1)
	p.run()
}
