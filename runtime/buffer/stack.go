package buffer

import (
	"errors"
	"fmt"
)

// ExecStack is the enviroments stack used for function calls, operations etc
type ExecStack struct {
	data []interface{}
}

const iNITIAL_STACK_SIZE = 128

// NewExecStack initialises a new stack
func NewExecStack() *ExecStack {
	s := &ExecStack{
		data: make([]interface{}, 0, iNITIAL_STACK_SIZE),
	}

	return s
}

// Reset resets the stack top to 0
func (s *ExecStack) Reset(val interface{}) {
	s.data = s.data[0:0]
}

// Push pushes a value into the stack
func (s *ExecStack) Push(val interface{}) {
	s.data = append(s.data, val)
}

// Pop pops a value from the stack
func (s *ExecStack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("execution stack empty")
	}

	top := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return top, nil
}

// Dump dumps the stack in readable form
func (s *ExecStack) Dump() {
	fmt.Printf("=====STACK=====\n")
	for _, val := range s.data {
		fmt.Printf("%5v", val)
	}
	fmt.Println()
}
