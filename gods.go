package gods

import (
	"github.com/MehdiEidi/gods/bitset"
	"github.com/MehdiEidi/gods/deque"
	"github.com/MehdiEidi/gods/linkedlist/circularly"
	"github.com/MehdiEidi/gods/linkedlist/doubly"
	"github.com/MehdiEidi/gods/linkedlist/singly"
	"github.com/MehdiEidi/gods/queue"
	"github.com/MehdiEidi/gods/set"
	"github.com/MehdiEidi/gods/stack"
)

func NewSinglyLinkedList[T comparable]() *singly.SinglyLinkedList[T] {
	return singly.New[T]()
}

func NewDoublyLinkedList[T comparable]() *doubly.DoublyLinkedList[T] {
	return doubly.New[T]()
}

func NewCircularlyLinkedList[T comparable]() *circularly.CircularlyLinkedList[T] {
	return circularly.New[T]()
}

func NewSet[T comparable]() *set.Set[T] {
	return set.New[T]()
}

func NewDeque[T comparable]() deque.Deque[T] {
	return deque.NewLinkedDeque[T]()
}

func NewLinkedQueue[T comparable]() queue.Queue[T] {
	return queue.NewLinkedQueue[T]()
}

func NewCircularQueue[T comparable]() queue.Queue[T] {
	return queue.NewCircularQueue[T]()
}

func NewSliceQueue[T any]() queue.Queue[T] {
	return queue.NewSliceQueue[T]()
}

func NewLinkedStack[T comparable]() stack.Stack[T] {
	return stack.NewLinkedStack[T]()
}

func NewSliceStack[T any]() stack.Stack[T] {
	return stack.NewSliceStack[T]()
}

func NewBitset8() *bitset.Bitset8 {
	return bitset.New()
}
