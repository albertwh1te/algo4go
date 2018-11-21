package algo4go

// MaxHeap always has maxium at root
type MaxHeap struct {
	arr  []int
	size int
}

// get heap parent element
func (heap MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

// get heap left element
func (heap MaxHeap) left(i int) int {
	if (2*i + 1) < heap.size {
		return 2*i + 1
	}
	return 0
}

// get heap right element
func (heap MaxHeap) right(i int) int {
	if (2*i + 2) < heap.size {
		return 2*i + 2
	}
	return 0
}

// TODO:fix me !
// func (heap *MaxHeap) Insert(n int) {
// 	(*heap).arr = append((*heap).arr, n)
// 	(*heap).size++
// }

// MaxHeapify assumes that the binary trees rooted at LEFT and RIGHT are max-heaps,
// but that node at i might be smaller than its children.
func (heap *MaxHeap) MaxHeapify(i int) {
	var largest int
	r := heap.right(i)
	l := heap.left(i)
	for l > 0 {
		if heap.arr[l] > heap.arr[i] {
			largest = l
		} else if (r > 0) && (heap.arr[r] > heap.arr[i]) {
			largest = r
		}

		if largest == i {
			return
		}
		// keep go down
		swap(&heap.arr[i], &heap.arr[largest])
		i = largest
		r = heap.right(i)
		l = heap.left(i)
	}
}

// BuildMaxHeap goes through the remaining nodes of the tree
// and runs MaxHeapify on each one
func (heap *MaxHeap) BuildMaxHeap() {
	for i := heap.size - 1; i >= 0; i-- {
		heap.MaxHeapify(i)
	}
}

// HeapInsert keeps max heap and insert element into Maxheap
func (heap *MaxHeap) HeapInsert(n int) {
	heap.arr = append(heap.arr, n)
	heap.size++
	index := heap.size - 1
	parent := heap.parent(index)
	for heap.arr[index] > heap.arr[parent] {
		swap(&heap.arr[index], &heap.arr[parent])
		index = parent
		parent = heap.parent(index)
	}
}
