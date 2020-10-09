package main

type HeapData interface {
	LessOrEqual(target interface{}) bool
}
type Heap struct {
	datas []HeapData
}
type HeapEmptyError struct {
}

func NewHeap() *Heap {
	heap := &Heap{}
	heap.datas = make([]HeapData, 1)
	return heap
}
func (heap *Heap) Pop() (HeapData, error) {
	maxIndex := len(heap.datas) - 1
	if maxIndex == 0 {
		return nil, HeapEmptyError{}
	}
	tmp := heap.datas[1]
	heap.datas[1] = heap.datas[maxIndex]
	heap.datas[maxIndex] = tmp
	maxIndex = maxIndex - 1
	shiftMin(heap.datas, 1, maxIndex)
	data := heap.datas[maxIndex+1]
	heap.datas = heap.datas[:maxIndex+1]
	return data, nil
}
func (heap *Heap) Push(data HeapData) {
	heap.datas = append(heap.datas, data)
	parent := (len(heap.datas) - 1) / 2
	addedDataIdx := len(heap.datas) - 1
	for parent >= 1 {
		if heap.datas[parent].LessOrEqual(heap.datas[addedDataIdx]) {
			break
		} else {
			tmp := heap.datas[parent]
			heap.datas[parent] = heap.datas[addedDataIdx]
			heap.datas[addedDataIdx] = tmp
			addedDataIdx = parent
			parent = parent / 2
		}
	}
}
func (heap *Heap) Len() int {
	return len(heap.datas) - 1
}
func shiftMin(data []HeapData, parent int, maxIndex int) {
	for parent <= maxIndex {
		left := parent * 2
		right := left + 1
		tmpMin := 0
		if right <= maxIndex {
			if !data[right].LessOrEqual(data[left]) {
				tmpMin = left
			} else {
				tmpMin = right
			}
		} else if left <= maxIndex {
			tmpMin = left
		} else {
			break
		}
		if !data[parent].LessOrEqual(data[tmpMin]) {
			tmp := data[parent]
			data[parent] = data[tmpMin]
			data[tmpMin] = tmp
			parent = tmpMin
		} else {
			break
		}

	}
}
func (heap *Heap) init() {
	maxIndex := len(heap.datas) - 1
	firstParent := maxIndex / 2
	for firstParent != 0 {
		parent := firstParent
		shiftMin(heap.datas, parent, maxIndex)
		firstParent = firstParent - 1
	}
}
func (err HeapEmptyError) Error() string {
	return "heap is empty"
}
func (err HeapEmptyError) Is(target error) bool {
	_, ok := target.(HeapEmptyError)
	return ok
}
