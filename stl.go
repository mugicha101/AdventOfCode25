package main

import (
	"container/heap"
	"sort"
)

type Pair[T, U any] struct {
	First  T
	Second U
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type Set[T comparable] map[T]bool

func (s Set[T]) Has(x T) bool {
	_, has := s[x]
	return has
}

func (s Set[T]) Insert(x T) bool {
	if s.Has(x) {
		return false
	}
	s[x] = true
	return true
}

func (s Set[T]) Erase(x T) bool {
	if !s.Has(x) {
		return false
	}

	delete(s, x)
	return true
}

type MultiSet[T comparable] map[T]int

func (s MultiSet[T]) Has(x T) bool {
	_, has := s[x]
	return has
}

func (s MultiSet[T]) Count(x T) int {
	amt, has := s[x]
	if has {
		return amt
	} else {
		return 0
	}
}

func (s MultiSet[T]) Insert(x T) {
	s[x]++
}

func (s MultiSet[T]) Delete(x T) bool {
	amt := s.Count(x)
	if amt == 0 {
		return false
	}
	if amt == 1 {
		delete(s, x)
	} else {
		s[x] = amt - 1
	}
	return true
}

type PairList[T, U Ordered] []Pair[T, U]

func (p PairList[T, U]) Len() int      { return len(p) }
func (p PairList[T, U]) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList[T, U]) Less(i, j int) bool {
	return p[i].First < p[j].First || (p[i].First == p[j].First && p[i].Second < p[j].Second)
}
func (p PairList[T, U]) Sort() {
	sort.Sort(p)
}

type MinHeap[T Ordered] []T

func (h *MinHeap[T]) Len() int {
	return len(*h)
}

func (h *MinHeap[T]) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *MinHeap[T]) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MaxHeap[T Ordered] []T

func (h *MaxHeap[T]) Len() int {
	return len(*h)
}

func (h *MaxHeap[T]) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *MaxHeap[T]) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MinPriorityQueue[T Ordered] MinHeap[T]

func (pq *MinPriorityQueue[T]) Push(x T) {
	heap.Push((*MinHeap[T])(pq), x)
}

func (pq *MinPriorityQueue[T]) Pop() T {
	return heap.Pop((*MinHeap[T])(pq)).(T)
}

func (pq *MinPriorityQueue[T]) Top() T {
	return (*pq)[0]
}

type MaxPriorityQueue[T Ordered] MaxHeap[T]

func (pq *MaxPriorityQueue[T]) Push(x T) {
	heap.Push((*MaxHeap[T])(pq), x)
}

func (pq *MaxPriorityQueue[T]) Pop() T {
	return heap.Pop((*MaxHeap[T])(pq)).(T)
}

func (pq *MaxPriorityQueue[T]) Top() T {
	return (*pq)[0]
}
