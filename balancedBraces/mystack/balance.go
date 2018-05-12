package mystack

import stack "github.com/golang-collections/collections/stack"

type MatchMap map[string]string

func NewBalanceChecker(m map[string]string) MatchMap {
	return MatchMap(m)
}

func (mm MatchMap) IsBalanced(s string) bool {
	// the empty string is balanced
	if len(s) == 0 {
		return true
	}

	// odd-length strings can never be balanced
	if len(s)%2 == 1 {
		return false
	}

	stack := stack.New()

	for i := 0; i < len(s); i++ {
		current := string(s[i])

		if stack.Len() == 0 {
			if _, isOpener := mm[current]; !isOpener {
				return false
			}

			stack.Push(current)
			continue
		}

		// adding an opener never makes something balanceable become un-balanceable
		if _, isOpener := mm[current]; isOpener {
			stack.Push(current)
			continue
		}

		// current is a closer
		// make sure that the pattern it makes with the previous char is not (}, [), etc.
		// i.e. if the first char is an opener, that opener must be of the
		// same type/class of bracket as second

		previous := stack.Peek().(string) // TODO be careful

		if _, isOpener := mm[previous]; isOpener {
			if mm[previous] != current {
				return false
			}

			// we have {}, (), etc
			// the current is not yet on the stack; the previous should be popped off
			stack.Pop()
			continue
		}

		// then adding current makes for 2 closers in a row
		// that may still be balanceable
		stack.Push(current)
	}

	return stack.Len() == 0
}