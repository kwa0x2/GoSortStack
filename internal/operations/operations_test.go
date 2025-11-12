package operations

import (
	"testing"

	"github.com/kwa0x2/GoSortStack/internal/stack"
)

func TestSwapA(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)

	SwapA(ops)

	if ops.StackA.Value != 1 {
		t.Errorf("Expected top value 1, got %d", ops.StackA.Value)
	}
	if ops.StackA.Next.Value != 2 {
		t.Errorf("Expected second value 2, got %d", ops.StackA.Next.Value)
	}
}

func TestSwapB(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToB(1)
	ops.PushToB(2)

	SwapB(ops)

	if ops.StackB.Value != 1 {
		t.Errorf("Expected top value 1, got %d", ops.StackB.Value)
	}
}

func TestSwapBoth(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToB(3)
	ops.PushToB(4)

	SwapBoth(ops)

	if ops.StackA.Value != 1 {
		t.Errorf("Expected StackA top value 1, got %d", ops.StackA.Value)
	}
	if ops.StackB.Value != 3 {
		t.Errorf("Expected StackB top value 3, got %d", ops.StackB.Value)
	}
}

func TestPushA(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToB(1)
	ops.PushToB(2)

	PushA(ops)

	if ops.StackA.Value != 2 {
		t.Errorf("Expected StackA top value 2, got %d", ops.StackA.Value)
	}
	if stack.Size(ops.StackB) != 1 {
		t.Errorf("Expected StackB size 1, got %d", stack.Size(ops.StackB))
	}
}

func TestPushB(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)

	PushB(ops)

	if ops.StackB.Value != 2 {
		t.Errorf("Expected StackB top value 2, got %d", ops.StackB.Value)
	}
	if stack.Size(ops.StackA) != 1 {
		t.Errorf("Expected StackA size 1, got %d", stack.Size(ops.StackA))
	}
}

func TestRotateA(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	RotateA(ops)

	if ops.StackA.Value != 2 {
		t.Errorf("Expected top value 2, got %d", ops.StackA.Value)
	}
}

func TestRotateB(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToB(1)
	ops.PushToB(2)
	ops.PushToB(3)

	RotateB(ops)

	if ops.StackB.Value != 2 {
		t.Errorf("Expected top value 2, got %d", ops.StackB.Value)
	}
}

func TestRotateBoth(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)
	ops.PushToB(4)
	ops.PushToB(5)
	ops.PushToB(6)

	RotateBoth(ops)

	if ops.StackA.Value != 2 {
		t.Errorf("Expected StackA top value 2, got %d", ops.StackA.Value)
	}
	if ops.StackB.Value != 5 {
		t.Errorf("Expected StackB top value 5, got %d", ops.StackB.Value)
	}
}

func TestReverseRotateA(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	ReverseRotateA(ops)

	if ops.StackA.Value != 1 {
		t.Errorf("Expected top value 1, got %d", ops.StackA.Value)
	}
}

func TestReverseRotateB(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToB(1)
	ops.PushToB(2)
	ops.PushToB(3)

	ReverseRotateB(ops)

	if ops.StackB.Value != 1 {
		t.Errorf("Expected top value 1, got %d", ops.StackB.Value)
	}
}

func TestReverseRotateBoth(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)
	ops.PushToB(4)
	ops.PushToB(5)
	ops.PushToB(6)

	ReverseRotateBoth(ops)

	if ops.StackA.Value != 1 {
		t.Errorf("Expected StackA top value 1, got %d", ops.StackA.Value)
	}
	if ops.StackB.Value != 4 {
		t.Errorf("Expected StackB top value 4, got %d", ops.StackB.Value)
	}
}

func TestEmptyStackOperations(t *testing.T) {
	ops := stack.NewOperations(true)

	SwapA(ops)
	SwapB(ops)
	PushA(ops)
	PushB(ops)
}
