package main

import "sort"

type Pair[T, U any] struct {
	First  T
	Second U
}

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type set[T comparable] map[T]bool

func (s set[T]) Has(x T) bool {
	_, has := s[x]
	return has
}

func (s set[T]) Insert(x T) bool {
	if s.Has(x) {
		return false
	}
	s[x] = true
	return true
}

func (s set[T]) Erase(x T) bool {
	if !s.Has(x) {
		return false
	}

	delete(s, x)
	return true
}

type multiset[T comparable] map[T]int

func (s multiset[T]) Has(x T) bool {
	_, has := s[x]
	return has
}

func (s multiset[T]) Count(x T) int {
	amt, has := s[x]
	if has {
		return amt
	} else {
		return 0
	}
}

func (s multiset[T]) Insert(x T) {
	s[x]++
}

func (s multiset[T]) Delete(x T) bool {
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
