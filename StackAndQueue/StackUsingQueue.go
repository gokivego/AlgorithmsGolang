/*
Implemet a Stack using Queues

Method 1: Use two Queues to implement a stack. PUSH:O()	POP:O()

Method 2: Use two Queues to implement a stack. PUSH:O()	POP:O()

"github.com/golang-collections/collections/stack"
type Stack
func New() *Stack
func (this *Stack) Len() int
func (this *Stack) Peek() interface{}
func (this *Stack) Pop() interface{}
func (this *Stack) Push(value interface{})

"github.com/golang-collections/go-datastructures/queue"
type Queue
func New(hint int64) *Queue
func (q *Queue) Dispose()
func (q *Queue) Disposed() bool
func (q *Queue) Empty() bool
func (q *Queue) Get(number int64) ([]interface{}, error)
func (q *Queue) Len() int64
func (q *Queue) Put(items ...interface{}) error
func (q *Queue) TakeUntil(checker func(item interface{}) bool) ([]interface{}, error)

*/

package main

import ("github.com/golang-collections/go-datastructures/queue"
		"github.com/golang-collections/collections/stack"
)

type StackUsingQueue struct {
	// Initialized
}

func (*)

func main() {

}