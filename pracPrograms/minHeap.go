package pracPrograms

import "fmt"

type MinHeap struct {
	items []int
}

func (h *MinHeap) Insert(val int) {
	h.items = append(h.items, val)
	h.heapifyUp(len(h.items) - 1)
}

func (h *MinHeap) heapifyUp(i int) {
	for i > 0 && h.items[i] < h.items[(i-1)/2] {
		parent := (i - 1) / 2
		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		i = parent
	}
}

func (h *MinHeap) Remove() {
	if len(h.items) == 0 {
		panic("Heap is empty")
	}
	removed := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	last := len(h.items) - 1
	h.items = h.items[:last]

	h.heapifyDown()
	fmt.Println("Removed ", removed)
}

func (h *MinHeap) heapifyDown() {
	start := 0
	lastIndex := len(h.items) - 1

	for {

		left := start*2 + 1
		right := start*2 + 2
		min := start

		if left <= lastIndex && h.items[left] < h.items[min] {
			min = left
		}
		if right <= lastIndex && h.items[right] < h.items[min] {
			min = right
		}

		if min != start {
			h.items[start], h.items[min] = h.items[min], h.items[start]
			start = min
		} else {
			break
		}
	}
}

func (h *MinHeap) Print() {
	fmt.Println(h.items)
}
