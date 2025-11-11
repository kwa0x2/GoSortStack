package checker

import "github.com/kwa0x2/GoSortStack/internal/stack"

func IsSorted(head *stack.Node) bool {
	if head == nil {
		return true 
	}

	current := head
	for current.Next != nil {
		if current.Value > current.Next.Value {
			return false
		}
		current = current.Next
	}

	return true
}

func IsSortedWithEmptyB(ops *stack.Operations) bool {
	return IsSorted(ops.StackA) && ops.StackB == nil
}
