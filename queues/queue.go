package queues

import "fmt"

type Data struct {
	ArrivalTime   float64
	DepartureTime float64
}

type Node struct {
	data Data
	next *Node
}

type Queue struct {
	currSize int
	front    *Node
	rear     *Node
}

func InitQueue() *Queue {
	return &Queue{0, nil, nil}
}

func GetQueueSize(q *Queue) int {
	return q.currSize
}

func EnQueue(q *Queue, d Data) {
	new := &Node{d, nil}

	if q.front == nil {
		q.front = new
	} else {
		q.rear.next = new
	}
	q.rear = new

	q.currSize++
}

func DeQueue(q *Queue) Data {
	data := Data{0, 0}

	if q.currSize > 0 {
		tmp := q.front
		data = tmp.data
		q.front = q.front.next
		q.currSize--
		if q.currSize == 0 {
			q.rear = nil
		}
	}

	return data
}

func PrintQueue(q *Queue) {
	for n := q.front; n != nil; n = n.next {
		fmt.Printf("arrival: %v departure: %v\n", n.data.ArrivalTime, n.data.DepartureTime)
	}
}

func FreeQueue(q *Queue) {
	for q.front != nil {
		tmp := q.front
		q.front = q.front.next
		tmp.next = nil
	}
	q.rear = nil
	q.currSize = 0
}
