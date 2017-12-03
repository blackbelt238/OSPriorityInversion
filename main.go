package main

// runs the simulation
func main() {
	p := createProcessor()

	p.addJob(1, t3)
	p.addJob(2, t1)
	p.addJob(6, t2)
	p.run()
}
