package algorithm

import (
	"github.com/kwa0x2/GoSortStack/internal/operations"
	"github.com/kwa0x2/GoSortStack/internal/stack"
)

func rotateTypeAB(a, b *stack.Node) int {
	tmp := a
	minCost := caseRraRrb(a, b, a.Value)

	for tmp != nil {
		if cost := caseRaRb(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		if cost := caseRraRrb(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		if cost := caseRaRrb(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		if cost := caseRraRb(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		tmp = tmp.Next
	}

	return minCost
}

func rotateTypeBA(a, b *stack.Node) int {
	tmp := b
	minCost := caseRraRrbA(a, b, b.Value)

	for tmp != nil {
		if cost := caseRaRbA(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		if cost := caseRraRrbA(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		if cost := caseRaRrbA(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		if cost := caseRraRbA(a, b, tmp.Value); minCost > cost {
			minCost = cost
		}
		tmp = tmp.Next
	}

	return minCost
}

func sortBTill3(ops *stack.Operations) {
	for stack.Size(ops.StackA) > 3 && !stack.IsSorted(ops.StackA) {
		tmp := ops.StackA
		minCost := rotateTypeAB(ops.StackA, ops.StackB)

		for minCost >= 0 {
			if minCost == caseRaRb(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRaRb(ops, tmp.Value, 'a')
			} else if minCost == caseRraRrb(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRraRrb(ops, tmp.Value, 'a')
			} else if minCost == caseRaRrb(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRaRrb(ops, tmp.Value, 'a')
			} else if minCost == caseRraRb(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRraRb(ops, tmp.Value, 'a')
			} else {
				tmp = tmp.Next
			}
		}
	}
}

func sortB(ops *stack.Operations) {
	if stack.Size(ops.StackA) > 3 && !stack.IsSorted(ops.StackA) {
		operations.PushB(ops)
	}
	if stack.Size(ops.StackA) > 3 && !stack.IsSorted(ops.StackA) {
		operations.PushB(ops)
	}
	if stack.Size(ops.StackA) > 3 && !stack.IsSorted(ops.StackA) {
		sortBTill3(ops)
	}
	if !stack.IsSorted(ops.StackA) {
		SortThree(ops)
	}
}

func sortA(ops *stack.Operations) {
	for ops.StackB != nil {
		tmp := ops.StackB
		minCost := rotateTypeBA(ops.StackA, ops.StackB)

		for minCost >= 0 {
			if minCost == caseRaRbA(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRaRb(ops, tmp.Value, 'b')
			} else if minCost == caseRaRrbA(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRaRrb(ops, tmp.Value, 'b')
			} else if minCost == caseRraRrbA(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRraRrb(ops, tmp.Value, 'b')
			} else if minCost == caseRraRbA(ops.StackA, ops.StackB, tmp.Value) {
				minCost = applyRraRb(ops, tmp.Value, 'b')
			} else {
				tmp = tmp.Next
			}
		}
	}
}

func Sort(ops *stack.Operations) {
	if stack.IsSorted(ops.StackA) {
		return
	}

	size := stack.Size(ops.StackA)

	if size == 2 {
		if ops.StackA.Value > ops.StackA.Next.Value {
			operations.SwapA(ops)
		}
		return
	}

	sortB(ops)

	sortA(ops)

	minValue := stack.Min(ops.StackA)
	minIndex := stack.FindIndex(ops.StackA, minValue)
	stackSize := stack.Size(ops.StackA)

	if minIndex < stackSize-minIndex {
		for ops.StackA.Value != minValue {
			operations.RotateA(ops)
		}
	} else {
		for ops.StackA.Value != minValue {
			operations.ReverseRotateA(ops)
		}
	}
}
