package heap

type Element[T any] struct {
	Priority int
	Key      T
}

type Heap[T any] []Element[T]

func New[T any]() Heap[T] {
	return Heap[T]{}
}

func (h Heap[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap[T]) Push(key T, priority int) {
	*h = append(*h, Element[T]{Key: key, Priority: priority})
	i := len(*h) - 1

	for i > 0 && priority > (*h)[(i-1)/2].Priority {
		h.Swap(i, (i-1)/2)
		i = (i - 1) / 2
	}
}

func (h *Heap[T]) Pop() Element[T] {
	n := len(*h)
	if n == 0 {
		return Element[T]{}
	}

	max := (*h)[0]
	h.Swap(0, n-1)
	(*h) = (*h)[:n-1]

	h.heapify()

	return max
}

func (h Heap[T]) heapify() {
	i := 0
	n := len(h)
	for i*2+1 < n {
		left := i*2 + 1
		right := left + 1
		swapIdx := left

		if right < n && h[left].Priority < h[right].Priority {
			swapIdx = right
		}

		if h[swapIdx].Priority > h[i].Priority {
			h.Swap(i, swapIdx)
			i = swapIdx

		} else {
			break
		}
	}
}
