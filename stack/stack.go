package stack

import "errors"

var (
	// ErrStackEmpty is returned if an operation is done on a stack
	// which requires a stack to have elements (Pop, Top) but the
	// stack has no elements.
	ErrStackEmpty = errors.New("Stack: stack is empty")
)

// Stack is a basic FIFO stack implementation
type Stack []interface{}

// New returns a new stack
func New() *Stack {
	return &Stack{}
}

// Push pushes item x onto the top of the stack
func (s *Stack) Push(x interface{}) {
	*s = append(*s, x)
}

// Pop returns the item at the top of the stack
func (s *Stack) Pop() (interface{}, error) {
	if len(*s) == 0 {
		return nil, ErrStackEmpty
	}
	lsti := len(*s) - 1
	tmp := (*s)[lsti]
	*s = (*s)[:lsti]
	return tmp, nil
}

// Top returns but doesn't remove the item at the top of the stack
func (s *Stack) Top() (interface{}, error) {
	if len(*s) == 0 {
		return nil, ErrStackEmpty
	}
	return (*s)[len(*s)-1], nil
}

// Len returns the number of items in the stack
func (s *Stack) Len() int {
	return len(*s)
}
