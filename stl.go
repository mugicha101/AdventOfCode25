package main

import (
	"container/heap"
	"sort"

	"github.com/tidwall/btree"
)

type Pair[T, U any] struct {
	A T
	B U
}

type OrderedPair[T Ordered, U Ordered] Pair[T, U]

func (a *OrderedPair[T, U]) Less(b *OrderedPair[T, U]) bool {
	return a.A < b.A || (a.A == b.A && a.B < b.B)
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// hashset without duplicates

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

// hashset with duplicates allowed

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

// sortable list of pairs

type OrderedPairList[T, U Ordered] []OrderedPair[T, U]

func (p OrderedPairList[T, U]) Len() int      { return len(p) }
func (p OrderedPairList[T, U]) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p OrderedPairList[T, U]) Less(i, j int) bool {
	return p[i].Less(&p[j])
}
func (p OrderedPairList[T, U]) Sort() {
	sort.Sort(p)
}

// min heap interface implementation

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

// max heap interface implementation

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

// converts min heap interface to methods

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

// converts max heap interface to methods

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

// double ended queue

type Deque[T any] []T

func (q *Deque[T]) PushBack(x T) {
	*q = append(*q, x)
}

func (q *Deque[T]) PopBack() T {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}

func (q *Deque[T]) Front() T {
	return (*q)[0]
}

func (q *Deque[T]) PushFront(x T) {
	*q = append([]T{x}, *q...)
}

func (q *Deque[T]) PopFront() T {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q *Deque[T]) Back() T {
	return (*q)[len(*q)-1]
}

// fifo queue

type Queue[T any] Deque[T]

func (q *Queue[T]) Push(x T) {
	(*Deque[T])(q).PushBack(x)
}

func (q *Queue[T]) Pop() T {
	return (*Deque[T])(q).PopFront()
}

func (q *Queue[T]) Front() T {
	return (*Deque[T])(q).Front()
}

func (q *Queue[T]) Back() T {
	return (*Deque[T])(q).Back()
}

// lifo stack

type Stack[T any] Deque[T]

func (s *Stack[T]) Push(x T) {
	(*Deque[T])(s).PushBack(x)
}

func (s *Stack[T]) Pop() T {
	return (*Deque[T])(s).PopBack()
}

func (s *Stack[T]) Top() T {
	return (*Deque[T])(s).Back()
}

// ordered map
type OrderedMap[K Ordered, V any] btree.Map[K, V]

func NewOrderedMap[K Ordered, V any]() *OrderedMap[K, V] {
	return (*OrderedMap[K, V])(btree.NewMap[K, V](2))
}

func (m *OrderedMap[K, V]) Get(key K) V {
	v, found := (*btree.Map[K, V])(m).Get(key)
	if !found {
		var zero V
		return zero
	}
	return v
}

func (m *OrderedMap[K, V]) Set(key K, value V) {
	(*btree.Map[K, V])(m).Set(key, value)
}

func (m *OrderedMap[K, V]) Erase(key K) bool {
	_, found := (*btree.Map[K, V])(m).Delete(key)
	return found
}

func (m *OrderedMap[K, V]) Size() int {
	return (*btree.Map[K, V])(m).Len()
}

func (m *OrderedMap[K, V]) HasKey(key K) bool {
	_, found := (*btree.Map[K, V])(m).Get(key)
	return found
}

func (m *OrderedMap[K, V]) MinKey() K {
	if m.Size() == 0 {
		panic("MinKey called on empty OrderedMap")
	}
	key, _, _ := (*btree.Map[K, V])(m).Min()
	return key
}

func (m *OrderedMap[K, V]) MaxKey() K {
	if m.Size() == 0 {
		panic("MaxKey called on empty OrderedMap")
	}
	key, _, _ := (*btree.Map[K, V])(m).Max()
	return key
}

// ordered set using B-tree
type OrderedSet[T Ordered] OrderedMap[T, struct{}]

func NewOrderedSet[T Ordered]() *OrderedSet[T] {
	return (*OrderedSet[T])(NewOrderedMap[T, struct{}]())
}

func (s *OrderedSet[T]) Has(x T) bool {
	return (*OrderedMap[T, struct{}])(s).HasKey(x)
}

func (s *OrderedSet[T]) Insert(x T) bool {
	sizeBefore := s.Size()
	(*OrderedMap[T, struct{}])(s).Set(x, struct{}{})
	return s.Size() != sizeBefore
}

func (s *OrderedSet[T]) Erase(x T) bool {
	return (*OrderedMap[T, struct{}])(s).Erase(x)
}

func (s *OrderedSet[T]) Size() int {
	return (*OrderedMap[T, struct{}])(s).Size()
}

func (s *OrderedSet[T]) Min() T {
	return (*OrderedMap[T, struct{}])(s).MinKey()
}

func (s *OrderedSet[T]) Max() T {
	return (*OrderedMap[T, struct{}])(s).MaxKey()
}

// ordered multiset
type OrderedMultiSet[T Ordered] OrderedMap[T, int]

func NewOrderedMultiSet[T Ordered]() *OrderedMultiSet[T] {
	return (*OrderedMultiSet[T])(NewOrderedMap[T, int]())
}

func (s *OrderedMultiSet[T]) Count(x T) int {
	return (*OrderedMap[T, int])(s).Get(x)
}

func (s *OrderedMultiSet[T]) Insert(x T) {
	(*OrderedMap[T, int])(s).Set(x, s.Count(x)+1)
}

func (s *OrderedMultiSet[T]) Erase(x T) bool {
	count := s.Count(x)
	if count == 0 {
		return false
	}

	if count == 1 {
		(*OrderedMap[T, int])(s).Erase(x)
	} else {
		(*OrderedMap[T, int])(s).Set(x, count-1)
	}

	return true
}

func (s *OrderedMultiSet[T]) Size() int {
	return (*OrderedMap[T, int])(s).Size()
}

func (s *OrderedMultiSet[T]) Min() T {
	return (*OrderedMap[T, int])(s).MinKey()
}

func (s *OrderedMultiSet[T]) Max() T {
	return (*OrderedMap[T, int])(s).MaxKey()
}
