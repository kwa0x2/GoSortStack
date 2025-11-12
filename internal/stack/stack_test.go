package stack

import "testing"

func TestNewOperations(t *testing.T) {
	ops := NewOperations(false)
	if ops.StackA != nil {
		t.Error("Expected StackA to be nil")
	}
	if ops.StackB != nil {
		t.Error("Expected StackB to be nil")
	}
	if ops.Silent {
		t.Error("Expected Silent to be false")
	}
}

func TestPushToA(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	if ops.StackA.Value != 3 {
		t.Errorf("Expected top value 3, got %d", ops.StackA.Value)
	}
}

func TestPushToB(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToB(10)
	ops.PushToB(20)

	if ops.StackB.Value != 20 {
		t.Errorf("Expected top value 20, got %d", ops.StackB.Value)
	}
}

func TestSwapStack(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(1)
	ops.PushToA(2)

	SwapStack(&ops.StackA)

	if ops.StackA.Value != 1 {
		t.Errorf("Expected top value 1 after swap, got %d", ops.StackA.Value)
	}
	if ops.StackA.Next.Value != 2 {
		t.Errorf("Expected second value 2 after swap, got %d", ops.StackA.Next.Value)
	}
}

func TestRotateStack(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	RotateStack(&ops.StackA)

	if ops.StackA.Value != 2 {
		t.Errorf("Expected top value 2 after rotate, got %d", ops.StackA.Value)
	}
}

func TestReverseRotateStack(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	ReverseRotateStack(&ops.StackA)

	if ops.StackA.Value != 1 {
		t.Errorf("Expected top value 1 after reverse rotate, got %d", ops.StackA.Value)
	}
}

func TestPushStack(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(1)
	ops.PushToA(2)

	PushStack(&ops.StackB, &ops.StackA)

	if ops.StackB.Value != 2 {
		t.Errorf("Expected StackB top value 2, got %d", ops.StackB.Value)
	}
	if ops.StackA.Value != 1 {
		t.Errorf("Expected StackA top value 1, got %d", ops.StackA.Value)
	}
}

func TestSize(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	size := Size(ops.StackA)
	if size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}
}

func TestMin(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(5)
	ops.PushToA(2)
	ops.PushToA(8)
	ops.PushToA(1)

	min := Min(ops.StackA)
	if min != 1 {
		t.Errorf("Expected min 1, got %d", min)
	}
}

func TestMax(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(5)
	ops.PushToA(2)
	ops.PushToA(8)
	ops.PushToA(1)

	max := Max(ops.StackA)
	if max != 8 {
		t.Errorf("Expected max 8, got %d", max)
	}
}

func TestFindIndex(t *testing.T) {
	ops := NewOperations(false)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	index := FindIndex(ops.StackA, 2)
	if index != 1 {
		t.Errorf("Expected index 1, got %d", index)
	}

	index = FindIndex(ops.StackA, 99)
	if index != -1 {
		t.Errorf("Expected index -1 for non-existent value, got %d", index)
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		expected bool
	}{
		{"Empty stack", []int{}, true},
		{"Single element", []int{1}, true},
		{"Sorted ascending", []int{1, 2, 3, 4, 5}, true},
		{"Not sorted", []int{3, 1, 2}, false},
		{"Reverse sorted", []int{5, 4, 3, 2, 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ops := NewOperations(false)
			for i := len(tt.values) - 1; i >= 0; i-- {
				ops.PushToA(tt.values[i])
			}

			result := IsSorted(ops.StackA)
			if result != tt.expected {
				t.Errorf("IsSorted() = %v, want %v", result, tt.expected)
			}
		})
	}
}
