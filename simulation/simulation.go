package simulation

import (
	"github.com/haraldkoch/golang-play/queues"
	"fmt"
	"math"
)

type Simulation struct {
	runTime        float64
	arrivalRate    float64
	processingTime float64
}

func getNextArrival(rate float64) float64 {
	interval := 1.0 / rate;
	return math.Abs(GenerateGaussianNoise(interval, interval))
}

func InitSimulation(runtime, rate, processingTime float64) *Simulation {
	return &Simulation{runtime, rate, processingTime}
}

func RunSimulation(sim *Simulation) {
	var currTime float64 = 0
	var nextArrivalTime float64 =  getNextArrival(sim.arrivalRate)
	var nextDeliveryTime float64 = nextArrivalTime + sim.processingTime

	routerQueue := queues.InitQueue()
	eventQueue := queues.InitQueue()

	for currTime < sim.runTime {
		if (nextDeliveryTime < nextArrivalTime) {
			// update clock
			currTime = nextDeliveryTime

			if queues.GetQueueSize(routerQueue) > 0 {
				// deliver a packet from the queue
				d := queues.DeQueue(routerQueue)
				d.DepartureTime = currTime
				queues.EnQueue(eventQueue, d)

				// figure out next delivery time
				nextDeliveryTime = currTime + sim.processingTime
				fmt.Printf("Time: %8.4f D q: %2d a: %8.4f d: %8.4f next-d: %8.4f\n", currTime, queues.GetQueueSize(routerQueue), d.ArrivalTime, d.DepartureTime - d.ArrivalTime, nextDeliveryTime)
			} else {
				nextDeliveryTime = nextArrivalTime + sim.processingTime
				fmt.Printf("Time: %8.4f D q: %2d --------- nil --------- next-d: %8.4f\n", currTime, queues.GetQueueSize(routerQueue), nextDeliveryTime)
			}
		} else {
			// update clock
			currTime = nextArrivalTime

			// enqueue a new packet
			d := queues.Data{currTime, 0}
			queues.EnQueue(routerQueue, d);
			if queues.GetQueueSize(routerQueue) == 1 {
				// first packet, reset deliveryTime
				// this should be identical to setting it when the queue is empty above
				nextDeliveryTime = currTime + sim.processingTime
			}

			// next ArrivalTime
			nextArrivalTime = currTime + getNextArrival(sim.arrivalRate)

			fmt.Printf("Time: %8.4f A q: %2d                         next-a: %8.4f\n", currTime, queues.GetQueueSize(routerQueue), nextArrivalTime)
		}
	}
}