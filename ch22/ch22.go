package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val any) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() any {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val any) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() any {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

type Element struct {
	Value interface{}
	Next  *Element
	Prev  *Element
}

func main() {

	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}
	v1 := queue.Pop()
	for v1 != nil {
		fmt.Printf("%v -> ", v1)
		v1 = queue.Pop()
	}

	stack := NewStack()

	for i := 1; i < 5; i++ {
		stack.Push(i)
	}

}
