package main

import "github.com/haraldkoch/golang-play/simulation"

func main() {
	var sim *simulation.Simulation
	sim = simulation.InitSimulation(100.0, 2, 0.25)

	simulation.RunSimulation(sim)
}
