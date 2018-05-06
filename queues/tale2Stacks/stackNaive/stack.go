package stackNaive

// implements stack
type stack []int

func New() *stack {
	s := make(stack, 0)
	return &s
}

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

func (s *stack) Pop() int {
	if len(*s) == 0 {
		panic("Attempted to pop from empty stack")
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) Peek() int {
	if len(*s) == 0 {
		panic("Attempted to peek on empty stack")
	}

	return (*s)[len(*s)-1]
}
