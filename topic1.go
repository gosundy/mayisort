package main

import "errors"

type PriorityQueue struct {
	heap *Heap
}
type PriorityQueueData struct {
	priority int
}

type QueueEmptyError struct {
}
type Integer int

func NewPriorityQueue() *PriorityQueue {
	queue := &PriorityQueue{}
	queue.heap = NewHeap()
	return queue
}
func (queue *PriorityQueue) Pop() (PriorityQueueData, error) {
	data, err := queue.heap.Pop()
	if err != nil {
		if errors.Is(err, HeapEmptyError{}) {
			return PriorityQueueData{}, QueueEmptyError{}
		} else {
			return PriorityQueueData{}, err
		}
	}

	return data.(PriorityQueueData), nil
}
func (queue *PriorityQueue) Push(data HeapData) {
	queue.heap.Push(data)
}
func (queue *PriorityQueue) Len() int {
	return queue.heap.Len()
}

func (a PriorityQueueData) LessOrEqual(target interface{}) bool {
	b := target.(PriorityQueueData)
	if a.priority <= b.priority {
		return true
	} else {
		return false
	}
}
func (a Integer) LessOrEqual(target interface{}) bool {
	b := target.(Integer)
	if a < b {
		return true
	} else {
		return false
	}
}

func (err QueueEmptyError) Error() string {
	return "queue is empty"
}
func (err QueueEmptyError) Is(target error) bool {
	_, ok := target.(QueueEmptyError)
	return ok
}
