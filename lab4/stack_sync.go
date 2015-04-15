// +build !solution

// Leave an empty line above this comment.
package lab4

import "sync"

type SafeStack struct {
	top  *Element
	size int
	mut  sync.Mutex
}

func (ss *SafeStack) Len() int {
	return ss.size
}
func (ss *SafeStack) Push(value interface{}) {
	ss.mut.Lock()
	defer ss.mut.Unlock()
	ss.top = &Element{value, ss.top}
	ss.size++
}
func (ss *SafeStack) Pop() (value interface{}) {
	if ss.size > 0 {
		ss.mut.Lock()
		defer ss.mut.Unlock()
		value, ss.top = ss.top.value, ss.top.next
		ss.size--
		return
	}
	return nil
}
