package xbctnr

import "github.com/starryck/x-lib-go/source/core/toolkit/xbvalue"

type Queue[T any] struct {
	size int
	head *QueueNode[T]
	tail *QueueNode[T]
}

type QueueNode[T any] struct {
	next  *QueueNode[T]
	value T
}

func (queue *Queue[T]) Size() int {
	return queue.size
}

func (queue *Queue[T]) Push(value T) {
	tail := queue.tail
	next := &QueueNode[T]{value: value}
	if tail == nil {
		queue.head = next
		queue.tail = next
	} else {
		tail.next = next
		queue.tail = next
	}
	queue.size++
}

func (queue *Queue[T]) Pull() (T, bool) {
	head := queue.head
	if head == nil {
		return xbvalue.Zero[T](), false
	}
	next := head.next
	queue.head = next
	queue.size--
	if next == nil {
		queue.tail = nil
	}
	return head.value, true
}

func (queue *Queue[T]) Poll() (T, bool) {
	head := queue.head
	if head == nil {
		return xbvalue.Zero[T](), false
	}
	if queue.size > 1 {
		tail := queue.tail
		queue.head = head.next
		head.next = nil
		tail.next = head
		queue.tail = head
	}
	return head.value, true
}

func (queue *Queue[T]) Peek() (T, bool) {
	head := queue.head
	if head == nil {
		return xbvalue.Zero[T](), false
	}
	return head.value, true
}

func (queue *Queue[T]) Slice() []T {
	next := queue.head
	if next == nil {
		return nil
	}
	slice := make([]T, queue.size)
	for i := 0; i < queue.size; i++ {
		slice[i] = next.value
		next = next.next
	}
	return slice
}

func (queue *Queue[T]) Iterator() *QueueIterator[T] {
	next := queue.head
	return &QueueIterator[T]{next: next}
}

func (queue *Queue[T]) Clear() {
	queue.size = 0
	queue.head = nil
	queue.tail = nil
}

type QueueIterator[T any] struct {
	next *QueueNode[T]
}

func (iterator *QueueIterator[T]) Next() (T, bool) {
	next := iterator.next
	if next == nil {
		return xbvalue.Zero[T](), false
	}
	iterator.next = next.next
	return next.value, true
}

type Deque[T any] struct {
	size int
	head *DequeNode[T]
	tail *DequeNode[T]
}

type DequeNode[T any] struct {
	prev  *DequeNode[T]
	next  *DequeNode[T]
	value T
}

func (deque *Deque[T]) Size() int {
	return deque.size
}

func (deque *Deque[T]) Push(value T) {
	tail := deque.tail
	next := &DequeNode[T]{value: value}
	if tail == nil {
		deque.head = next
		deque.tail = next
	} else {
		tail.next = next
		next.prev = tail
		deque.tail = next
	}
	deque.size++
}

func (deque *Deque[T]) Pull() (T, bool) {
	head := deque.head
	if head == nil {
		return xbvalue.Zero[T](), false
	}
	next := head.next
	deque.head = next
	deque.size--
	if next == nil {
		deque.tail = nil
	} else {
		next.prev = nil
	}
	return head.value, true
}

func (deque *Deque[T]) Poll() (T, bool) {
	head := deque.head
	if head == nil {
		return xbvalue.Zero[T](), false
	}
	if deque.size > 1 {
		tail := deque.tail
		next := head.next
		deque.head = next
		next.prev = nil
		tail.next = head
		head.prev = tail
		head.next = nil
		deque.tail = head
	}
	return head.value, true
}

func (deque *Deque[T]) Peek() (T, bool) {
	head := deque.head
	if head == nil {
		return xbvalue.Zero[T](), false
	}
	return head.value, true
}

func (deque *Deque[T]) Slice() []T {
	next := deque.head
	if next == nil {
		return nil
	}
	slice := make([]T, deque.size)
	for i := 0; i < deque.size; i++ {
		slice[i] = next.value
		next = next.next
	}
	return slice
}

func (deque *Deque[T]) Iterator() *DequeIterator[T] {
	next := deque.head
	return &DequeIterator[T]{deque: deque, next: next}
}

func (deque *Deque[T]) RPush(value T) {
	head := deque.head
	prev := &DequeNode[T]{value: value}
	if head == nil {
		deque.head = prev
		deque.tail = prev
	} else {
		head.prev = prev
		prev.next = head
		deque.head = prev
	}
	deque.size++
}

func (deque *Deque[T]) RPull() (T, bool) {
	tail := deque.tail
	if tail == nil {
		return xbvalue.Zero[T](), false
	}
	prev := tail.prev
	deque.tail = prev
	deque.size--
	if prev == nil {
		deque.head = nil
	} else {
		prev.next = nil
	}
	return tail.value, true
}

func (deque *Deque[T]) RPoll() (T, bool) {
	tail := deque.tail
	if tail == nil {
		return xbvalue.Zero[T](), false
	}
	if deque.size > 1 {
		head := deque.head
		prev := tail.prev
		deque.tail = prev
		prev.next = nil
		head.prev = tail
		tail.next = head
		tail.prev = nil
		deque.head = tail
	}
	return tail.value, true
}

func (deque *Deque[T]) RPeek() (T, bool) {
	tail := deque.tail
	if tail == nil {
		return xbvalue.Zero[T](), false
	}
	return tail.value, true
}

func (deque *Deque[T]) RSlice() []T {
	next := deque.tail
	if next == nil {
		return nil
	}
	slice := make([]T, deque.size)
	for i := 0; i < deque.size; i++ {
		slice[i] = next.value
		next = next.prev
	}
	return slice
}

func (deque *Deque[T]) RIterator() *DequeReverseIterator[T] {
	next := deque.tail
	return &DequeReverseIterator[T]{deque: deque, next: next}
}

func (deque *Deque[T]) Drop(node *DequeNode[T]) (T, bool) {
	switch node {
	case deque.head:
		return deque.Pull()
	case deque.tail:
		return deque.RPull()
	default:
		node.prev.next = node.next
		node.next.prev = node.prev
		deque.size--
		return node.value, true
	}
}

func (deque *Deque[T]) Clear() {
	deque.size = 0
	deque.head = nil
	deque.tail = nil
}

type DequeIterator[T any] struct {
	deque *Deque[T]
	prev  *DequeNode[T]
	next  *DequeNode[T]
}

func (iterator *DequeIterator[T]) Drop() (T, bool) {
	prev := iterator.prev
	if prev == nil {
		return xbvalue.Zero[T](), false
	}
	return iterator.deque.Drop(prev)
}

func (iterator *DequeIterator[T]) Next() (T, bool) {
	next := iterator.next
	if next == nil {
		return xbvalue.Zero[T](), false
	}
	iterator.prev = next
	iterator.next = next.next
	return next.value, true
}

type DequeReverseIterator[T any] struct {
	deque *Deque[T]
	prev  *DequeNode[T]
	next  *DequeNode[T]
}

func (iterator *DequeReverseIterator[T]) Drop() (T, bool) {
	prev := iterator.prev
	if prev == nil {
		return xbvalue.Zero[T](), false
	}
	return iterator.deque.Drop(prev)
}

func (iterator *DequeReverseIterator[T]) Next() (T, bool) {
	next := iterator.next
	if next == nil {
		return xbvalue.Zero[T](), false
	}
	iterator.prev = next
	iterator.next = next.prev
	return next.value, true
}
