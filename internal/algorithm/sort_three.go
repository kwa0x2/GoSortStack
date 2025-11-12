package algorithm

import (
	"github.com/kwa0x2/GoSortStack/internal/checker"
	"github.com/kwa0x2/GoSortStack/internal/operations"
	"github.com/kwa0x2/GoSortStack/internal/stack"
)


func SortThree(ops *stack.Operations) {
	if ops.StackA == nil || ops.StackA.Next == nil {
		return 
	}

	if checker.IsSorted(ops.StackA) {
		return
	}

	min := stack.Min(ops.StackA)
	max := stack.Max(ops.StackA)

	if ops.StackA.Value == min {
		operations.ReverseRotateA(ops)
		operations.SwapA(ops)
	} else if ops.StackA.Value == max {
		operations.RotateA(ops)
		if !checker.IsSorted(ops.StackA) {
			operations.SwapA(ops)
		}
	} else {
		maxIndex := stack.FindIndex(ops.StackA, max)
		if maxIndex == 1 {
			operations.ReverseRotateA(ops)
		} else {
			operations.SwapA(ops)
		}
	}
}
