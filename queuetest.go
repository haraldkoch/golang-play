package main

import "fmt"
import "github.com/haraldkoch/golang-play/queues"

func main() {
	testq := queues.InitQueue()
	fmt.Printf("Empty queue\n")
	queues.PrintQueue(testq)

	queues.EnQueue(testq, queues.Data{0.8, 1.2})
	queues.EnQueue(testq, queues.Data{1.0, 2.2})
	queues.EnQueue(testq, queues.Data{1.9, 3.2})

	fmt.Printf("\nFull queue\n")
	queues.PrintQueue(testq)

	queues.DeQueue(testq)

	fmt.Printf("\nDekete 1\n")
	queues.PrintQueue(testq)

	queues.FreeQueue(testq)
}
