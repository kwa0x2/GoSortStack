package algorithm

import (
	"github.com/kwa0x2/GoSortStack/internal/checker"
	"github.com/kwa0x2/GoSortStack/internal/operations"
	"github.com/kwa0x2/GoSortStack/internal/stack"
)

func SortFive(ops *stack.Operations) {
	if checker.IsSorted(ops.StackA) {
		return
	}

	size := stack.Size(ops.StackA)

	pushCount := size - 3
	for i := 0; i < pushCount; i++ {
		min := stack.Min(ops.StackA)
		minIndex := stack.FindIndex(ops.StackA, min)
		stackSize := stack.Size(ops.StackA)

		if minIndex <= stackSize/2 {
			for j := 0; j < minIndex; j++ {
				operations.RotateA(ops)
			}
		} else {
			for j := 0; j < stackSize-minIndex; j++ {
				operations.ReverseRotateA(ops)
			}
		}

		operations.PushB(ops)
	}

	SortThree(ops)

	for ops.StackB != nil {
		operations.PushA(ops)
	}
}
