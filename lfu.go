package lfu

import (
	"github.com/erfanmomeniii/heap"
)

const DefaultSize int64 = 2000

// Lfu is an instantiation of the lfu.
type Lfu struct {
	values    map[any]any
	frequents map[any]int
	minHeap   *heap.MinHeap
	size      int64
}

// Get returns the value associated with the inputted key.
func (l *Lfu) Get(key any) any {
	if _, ok := l.values[key]; ok {
		l.frequents[key]++
		l.minHeap.Update(key, l.frequents[key])
	}

	return l.values[key]
}

// purge removes the least used element from cache.
func (l *Lfu) purge() {
	if len(l.values) >= int(l.size) {
		_, key := l.minHeap.GetMin()
		l.minHeap.Delete()
		delete(l.values, key)
		delete(l.frequents, key)
	}

	return
}

// Set adds or updates with inputted key and value.
func (l *Lfu) Set(key any, value any) {
	if _, ok := l.values[key]; ok {
		l.values[key] = value
		l.frequents[key]++
		l.minHeap.Update(key, l.frequents[key])
	} else {
		l.purge()
		l.values[key] = value
		l.frequents[key] = 1
		l.minHeap.Insert(l.frequents[key], key)
	}

	return
}

// Has checks whether the key exists or not.
func (l *Lfu) Has(key any) bool {
	_, ok := l.values[key]

	return ok
}

// New creates a new instance of a lfu.
func New(size ...int64) *Lfu {
	s := DefaultSize
	if len(size) > 0 {
		s = size[0]
	}

	return &Lfu{values: make(map[any]any), frequents: make(map[any]int), minHeap: heap.NewMin(), size: s}
}
