package heap

type Heap struct {
	arr []int
}

func (h *Heap) build() {
	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		heapify(h.arr, i)
	}
}

func New(arr []int) *Heap {
	heap := &Heap{arr: arr}
	heap.build()
	return heap
}

func heapify(arr []int, i int) {
	smallest := i
	l, r := 2*i+1, 2*i+2
	if l < len(arr) && arr[smallest] > arr[l] {
		smallest = l
	}
	if r < len(arr) && arr[smallest] > arr[r] {
		smallest = r
	}
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		heapify(arr, smallest)
	}
}
