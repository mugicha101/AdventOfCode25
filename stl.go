package main

import (
	"container/heap"
	"sort"
	"strconv"

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

// general utility functions

func stoi(s string) int {
	res, ok := strconv.Atoi(s)
	if ok != nil {
		panic(ok)
	}
	return res
}

func stoll(s string) int64 {
	res, ok := strconv.ParseInt(s, 10, 64)
	if ok != nil {
		panic(ok)
	}
	return int64(res)
}

type Numeric interface {
	~int | ~float64 | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~uintptr | ~float32
}

type Integral interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~uintptr
}

func abs[T Numeric](x T) T {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

// union find

type UnionFind struct {
	uf    []int
	size  []int
	comps int
}

func NewUnionFind(n int) *UnionFind {
	uf := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		uf[i] = i
		size[i] = 1
	}
	return &UnionFind{uf: uf, size: size, comps: n}
}

func (uf *UnionFind) Find(x int) int {
	if uf.uf[x] != x {
		uf.uf[x] = uf.Find(uf.uf[x])
	}
	return uf.uf[x]
}

func (uf *UnionFind) Merge(a, b int) bool {
	ra := uf.Find(a)
	rb := uf.Find(b)
	if ra == rb {
		return false
	}
	if uf.size[ra] < uf.size[rb] {
		ra, rb = rb, ra
	}
	uf.uf[rb] = ra
	uf.size[ra] += uf.size[rb]
	uf.size[rb] = 0
	uf.comps--
	return true
}

func (uf *UnionFind) CompSize(x int) int {
	return uf.size[uf.Find(x)]
}

func (uf *UnionFind) NumComps() int {
	return uf.comps
}

func (uf *UnionFind) Len() int {
	return len(uf.uf)
}

// sorting algorithms
// note: modifies input array but also returns it just in case

// insertion sort based on second value of pair
func ISort[T any](arr []Pair[T, int64]) []Pair[T, int64] {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1].B > arr[j].B; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

// insertion sort directly on values

func ISortT[T Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

// quicksort based on second value of pair
func QSort[T any](arr []Pair[T, int64]) []Pair[T, int64] {
	if len(arr) < 5 {
		return ISort(arr)
	}

	// choose pivot
	mid := (len(arr)) >> 1
	if arr[0].B > arr[mid].B {
		arr[0], arr[mid] = arr[mid], arr[0]
	}
	if arr[mid].B > arr[len(arr)-1].B {
		arr[mid], arr[len(arr)-1] = arr[len(arr)-1], arr[mid]
	}
	if arr[0].B > arr[mid].B {
		arr[0], arr[mid] = arr[mid], arr[0]
	}

	arr[mid], arr[len(arr)-1] = arr[len(arr)-1], arr[mid]
	l := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].B < arr[len(arr)-1].B {
			arr[i], arr[l] = arr[l], arr[i]
			l++
		}
	}
	arr[l], arr[len(arr)-1] = arr[len(arr)-1], arr[l]
	QSort(arr[:l])
	QSort(arr[l+1:])
	return arr
}

// quicksort directly on values

func QSortT[T Ordered](arr []T) []T {
	if len(arr) < 5 {
		return ISortT(arr)
	}

	// choose pivot
	mid := (len(arr)) >> 1
	if arr[0] > arr[mid] {
		arr[0], arr[mid] = arr[mid], arr[0]
	}
	if arr[mid] > arr[len(arr)-1] {
		arr[mid], arr[len(arr)-1] = arr[len(arr)-1], arr[mid]
	}
	if arr[0] > arr[mid] {
		arr[0], arr[mid] = arr[mid], arr[0]
	}

	arr[mid], arr[len(arr)-1] = arr[len(arr)-1], arr[mid]
	l := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[len(arr)-1] {
			arr[i], arr[l] = arr[l], arr[i]
			l++
		}
	}
	arr[l], arr[len(arr)-1] = arr[len(arr)-1], arr[l]
	QSortT(arr[:l])
	QSortT(arr[l+1:])
	return arr
}

// base 2^bits radix sort based on second value of pair
func RSort[T any](arr []Pair[T, int64], bits int) []Pair[T, int64] {
	p := int64(1) << bits
	buckets := make([][]Pair[T, int64], p)
	next := make([]int, p)
	for i := 0; i < int(p); i++ {
		buckets[i] = make([]Pair[T, int64], len(arr))
	}
	maxBits := 0
	for i := 0; i < len(arr); i++ {
		for arr[i].B > (int64(1) << maxBits) {
			maxBits++
		}
	}
	maxBits += (maxBits & 1)
	for shift := 0; shift < maxBits; shift += 2 {
		for i := 0; i < len(arr); i++ {
			bit := (arr[i].B >> shift) & (p - 1)
			buckets[bit][next[bit]] = arr[i]
			next[bit]++
		}
		k := 0
		for i := 0; i < int(p); i++ {
			for j := 0; j < next[i]; j++ {
				arr[k] = buckets[i][j]
				k++
			}
		}
		for i := 0; i < int(p); i++ {
			next[i] = 0
		}
	}
	return arr
}

// coordinate compression helpers

// remove all duplicates from slice (must be grouped by equal values, shifts all duplicates to end and resizes slice to cut them off)
func Unique[T comparable](arr []T) []T {
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[j] != arr[i] {
			j++
			arr[j] = arr[i]
		}
	}
	return arr[:j+1]
}

// maps values in arr to their rank and returns sorted unique values + mapping from value to rank
func RankMap[T Ordered](arr []T) ([]T, map[T]int) {
	QSortT(arr)
	arr = Unique(arr)
	rank := map[T]int{}
	for i, v := range arr {
		rank[v] = i
	}
	return arr, rank
}

// init matrices
func Arr[T any](size int) []T {
	return make([]T, size)
}
func Mat[T any, N Integral](rows N, cols N) [][]T {
	arr := make([][]T, rows)
	for i := N(0); i < rows; i++ {
		arr[i] = make([]T, cols)
	}
	return arr
}
func Mat3D[T any, N Integral](x, y, z N) [][][]T {
	arr := make([][][]T, x)
	for i := N(0); i < x; i++ {
		arr[i] = Mat[T](y, z)
	}
	return arr
}
