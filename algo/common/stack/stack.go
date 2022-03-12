package stack

type Stack struct {
	vals []interface{}
}

func InitStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() interface{} {
	if len(s.vals) == 0 {
		return nil
	}
	end := len(s.vals) - 1
	result := s.vals[end]
	s.vals = s.vals[:end]

	return result
}

func (s *Stack) Push(val interface{}) {
	s.vals = append(s.vals, val)
}
