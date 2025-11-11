package algorithm

import (
	"github.com/kwa0x2/GoSortStack/internal/operations"
	"github.com/kwa0x2/GoSortStack/internal/stack"
)

func findPlaceB(stackB *stack.Node, nbrPush int) int {
	i := 1

	if stackB == nil {
		return 0
	}

	lastValue := stack.GetLast(stackB).Value

	if nbrPush > stackB.Value && nbrPush < lastValue {
		return 0
	}

	maxVal := stack.Max(stackB)
	minVal := stack.Min(stackB)
	if nbrPush > maxVal || nbrPush < minVal {
		return stack.FindIndex(stackB, maxVal)
	}

	tmp := stackB.Next
	current := stackB
	for current.Value < nbrPush || tmp.Value > nbrPush {
		current = current.Next
		tmp = current.Next
		i++
	}

	return i
}

func findPlaceA(stackA *stack.Node, nbrPush int) int {
	i := 1

	if stackA == nil {
		return 0
	}

	lastValue := stack.GetLast(stackA).Value

	if nbrPush < stackA.Value && nbrPush > lastValue {
		return 0
	}

	maxVal := stack.Max(stackA)
	minVal := stack.Min(stackA)
	if nbrPush > maxVal || nbrPush < minVal {
		return stack.FindIndex(stackA, minVal)
	}

	tmp := stackA.Next
	current := stackA
	for current.Value > nbrPush || tmp.Value < nbrPush {
		current = current.Next
		tmp = current.Next
		i++
	}

	return i
}

func caseRaRb(a, b *stack.Node, c int) int {
	placeB := findPlaceB(b, c)
	indexA := stack.FindIndex(a, c)

	if placeB < indexA {
		return indexA
	}
	return placeB
}

func caseRraRrb(a, b *stack.Node, c int) int {
	i := 0
	placeB := findPlaceB(b, c)
	indexA := stack.FindIndex(a, c)

	if placeB != 0 {
		i = stack.Size(b) - placeB
	}

	reverseIndexA := 0
	if indexA != 0 {
		reverseIndexA = stack.Size(a) - indexA
	}

	if i < reverseIndexA {
		return reverseIndexA
	}
	return i
}

func caseRraRb(a, b *stack.Node, c int) int {
	i := 0
	indexA := stack.FindIndex(a, c)

	if indexA != 0 {
		i = stack.Size(a) - indexA
	}

	return findPlaceB(b, c) + i
}

func caseRaRrb(a, b *stack.Node, c int) int {
	i := 0
	placeB := findPlaceB(b, c)

	if placeB != 0 {
		i = stack.Size(b) - placeB
	}

	return stack.FindIndex(a, c) + i
}

func caseRaRbA(a, b *stack.Node, c int) int {
	placeA := findPlaceA(a, c)
	indexB := stack.FindIndex(b, c)

	if placeA < indexB {
		return indexB
	}
	return placeA
}

func caseRraRrbA(a, b *stack.Node, c int) int {
	i := 0
	placeA := findPlaceA(a, c)
	indexB := stack.FindIndex(b, c)

	if placeA != 0 {
		i = stack.Size(a) - placeA
	}

	reverseIndexB := 0
	if indexB != 0 {
		reverseIndexB = stack.Size(b) - indexB
	}

	if i < reverseIndexB {
		return reverseIndexB
	}
	return i
}

func caseRaRrbA(a, b *stack.Node, c int) int {
	i := 0
	indexB := stack.FindIndex(b, c)

	if indexB != 0 {
		i = stack.Size(b) - indexB
	}

	return findPlaceA(a, c) + i
}

func caseRraRbA(a, b *stack.Node, c int) int {
	i := 0
	placeA := findPlaceA(a, c)

	if placeA != 0 {
		i = stack.Size(a) - placeA
	}

	return stack.FindIndex(b, c) + i
}

func applyRaRb(ops *stack.Operations, c int, s rune) int {
	if s == 'a' {
		for ops.StackA.Value != c && findPlaceB(ops.StackB, c) > 0 {
			operations.RotateBoth(ops)
		}
		for ops.StackA.Value != c {
			operations.RotateA(ops)
		}
		for findPlaceB(ops.StackB, c) > 0 {
			operations.RotateB(ops)
		}
		operations.PushB(ops)
	} else {
		for ops.StackB.Value != c && findPlaceA(ops.StackA, c) > 0 {
			operations.RotateBoth(ops)
		}
		for ops.StackB.Value != c {
			operations.RotateB(ops)
		}
		for findPlaceA(ops.StackA, c) > 0 {
			operations.RotateA(ops)
		}
		operations.PushA(ops)
	}
	return -1
}

func applyRraRrb(ops *stack.Operations, c int, s rune) int {
	if s == 'a' {
		for ops.StackA.Value != c && findPlaceB(ops.StackB, c) > 0 {
			operations.ReverseRotateBoth(ops)
		}
		for ops.StackA.Value != c {
			operations.ReverseRotateA(ops)
		}
		for findPlaceB(ops.StackB, c) > 0 {
			operations.ReverseRotateB(ops)
		}
		operations.PushB(ops)
	} else {
		for ops.StackB.Value != c && findPlaceA(ops.StackA, c) > 0 {
			operations.ReverseRotateBoth(ops)
		}
		for ops.StackB.Value != c {
			operations.ReverseRotateB(ops)
		}
		for findPlaceA(ops.StackA, c) > 0 {
			operations.ReverseRotateA(ops)
		}
		operations.PushA(ops)
	}
	return -1
}

func applyRraRb(ops *stack.Operations, c int, s rune) int {
	if s == 'a' {
		for ops.StackA.Value != c {
			operations.ReverseRotateA(ops)
		}
		for findPlaceB(ops.StackB, c) > 0 {
			operations.RotateB(ops)
		}
		operations.PushB(ops)
	} else {
		for findPlaceA(ops.StackA, c) > 0 {
			operations.ReverseRotateA(ops)
		}
		for ops.StackB.Value != c {
			operations.RotateB(ops)
		}
		operations.PushA(ops)
	}
	return -1
}

func applyRaRrb(ops *stack.Operations, c int, s rune) int {
	if s == 'a' {
		for ops.StackA.Value != c {
			operations.RotateA(ops)
		}
		for findPlaceB(ops.StackB, c) > 0 {
			operations.ReverseRotateB(ops)
		}
		operations.PushB(ops)
	} else {
		for findPlaceA(ops.StackA, c) > 0 {
			operations.RotateA(ops)
		}
		for ops.StackB.Value != c {
			operations.ReverseRotateB(ops)
		}
		operations.PushA(ops)
	}
	return -1
}
