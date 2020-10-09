package main

import (
	"errors"
	"math/rand"
	"sort"
	"testing"
)

//must more than 10
const TestDataSize = 10000

func init() {
	rand.Seed(123456)
}

func TestNewHeap(t *testing.T) {
	heap := NewHeap()
	_, err := heap.Pop()
	if !errors.Is(err, HeapEmptyError{}) {
		t.Fatalf("expect:no data, actual: having data")
	}
}
func TestHeapPushAndPop0(t *testing.T) {
	datas := rand.Perm(TestDataSize)
	heap := NewHeap()
	for i := 0; i < TestDataSize; i++ {
		heap.Push(Integer(datas[i]))
	}
	for i := 0; i < TestDataSize; i++ {
		data, err := heap.Pop()
		if errors.Is(err, HeapEmptyError{}) {
			t.Fatalf("expect:having data, actual: no  data")
		}
		if data != Integer(i) {
			t.Fatalf("expect:%d, acutal:%d", i, data)
		}
	}
	_, err := heap.Pop()
	if !errors.Is(err, HeapEmptyError{}) {
		t.Fatalf("expect:no data, actual: having data")
	}
}
func TestHeapPushAndPop1(t *testing.T) {
	datas := make([]int, 0)
	datas2 := make([]int, 0)
	for i := 0; i < TestDataSize; i++ {
		a := rand.Intn(TestDataSize * 100)
		datas = append(datas, a)
		datas2 = append(datas2, a)
	}
	sort.Ints(datas2)
	heap := NewHeap()
	for i := 0; i < TestDataSize; i++ {
		heap.Push(Integer(datas[i]))
	}
	for i := 0; i < TestDataSize; i++ {
		data, err := heap.Pop()
		if errors.Is(err, HeapEmptyError{}) {
			t.Fatalf("expect:having data, actual: no  data")
		}
		if data != Integer(datas2[i]) {
			t.Fatalf("expect:%d, acutal:%d", datas2[i], data)
		}
	}
	_, err := heap.Pop()
	if !errors.Is(err, HeapEmptyError{}) {
		t.Fatalf("expect:no data, actual: having data")
	}
}
func TestNewPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue()
	_, err := queue.Pop()
	if !errors.Is(err, QueueEmptyError{}) {
		t.Fatalf("expect:no data, actual: having data")
	}
}
func TestPriorityQueuePushAndPop0(t *testing.T) {
	datas := rand.Perm(TestDataSize)
	queue := NewPriorityQueue()
	for i := 0; i < TestDataSize; i++ {
		queue.Push(PriorityQueueData{priority: datas[i]})
	}
	for i := 0; i < TestDataSize; i++ {
		data, err := queue.Pop()
		if errors.Is(err, HeapEmptyError{}) {
			t.Fatalf("expect:having data, actual: no  data")
		}
		if data.priority != i {
			t.Fatalf("expect:%d, acutal:%d", i, data.priority)
		}
	}
	_, err := queue.Pop()
	if !errors.Is(err, QueueEmptyError{}) {
		t.Fatalf("expect:no data, actual: having data")
	}
}
func TestPriorityPushAndPop2(t *testing.T) {
	datas := make([]int, 0)
	datas2 := make([]int, 0)
	for i := 0; i < TestDataSize; i++ {
		a := rand.Intn(TestDataSize * 100)
		datas = append(datas, a)
		datas2 = append(datas2, a)
	}
	sort.Ints(datas2)
	queue := NewPriorityQueue()
	for i := 0; i < TestDataSize; i++ {
		queue.Push(PriorityQueueData{priority: datas[i]})
	}
	for i := 0; i < TestDataSize; i++ {
		data, err := queue.Pop()
		if errors.Is(err, QueueEmptyError{}) {
			t.Fatalf("expect:having data, actual: no  data")
		}
		if data.priority != datas2[i] {
			t.Fatalf("expect:%d, acutal:%d", datas2[i], data.priority)
		}
	}
	_, err := queue.Pop()
	if !errors.Is(err, QueueEmptyError{}) {
		t.Fatalf("expect:no data, actual: having data")
	}
}
func TestPriorityPushAndPop3(t *testing.T) {
	datas := make([]int, 0)
	datas2 := make([]int, 0)
	for i := 0; i < TestDataSize; i++ {
		a := rand.Intn(TestDataSize * 100)
		datas = append(datas, a)
		datas2 = append(datas2, a)
	}
	sort.Ints(datas2)
	queue := NewPriorityQueue()
	for i := 0; i < TestDataSize; i++ {
		queue.Push(PriorityQueueData{priority: datas[i]})
	}
	for i := 0; i < TestDataSize; i++ {
		data, err := queue.Pop()
		if errors.Is(err, QueueEmptyError{}) {
			t.Fatalf("expect:having data, actual: no  data")
		}
		if data.priority != datas2[i] {
			t.Fatalf("expect:%d, acutal:%d", datas2[i], data.priority)
		}
	}
	_, err := queue.Pop()
	if !errors.Is(err, QueueEmptyError{}) {
		t.Fatalf("expect:no data, actual: having data")
	}
	if queue.Len()!=0{
		t.Fatalf("expect:0, actual:%d",queue.Len())
	}
}
