package algorithm

import (
	"testing"

	"github.com/kwa0x2/GoSortStack/internal/stack"
)

func TestSortThreeAlreadySorted(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)

	SortThree(ops)

	if !stack.IsSorted(ops.StackA) {
		t.Error("Stack should be sorted")
	}
}

func TestSortThreeReverseSorted(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(3)
	ops.PushToA(2)
	ops.PushToA(1)

	SortThree(ops)

	if !stack.IsSorted(ops.StackA) {
		t.Error("Stack should be sorted")
	}
}

func TestSortThreeRandom(t *testing.T) {
	tests := [][]int{
		{2, 1, 3},
		{3, 1, 2},
		{1, 3, 2},
		{2, 3, 1},
	}

	for _, tt := range tests {
		ops := stack.NewOperations(true)
		for i := len(tt) - 1; i >= 0; i-- {
			ops.PushToA(tt[i])
		}

		SortThree(ops)

		if !stack.IsSorted(ops.StackA) {
			t.Errorf("Stack should be sorted for input %v", tt)
		}
	}
}

func TestSortFive(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)
	ops.PushToA(3)
	ops.PushToA(4)
	ops.PushToA(5)

	SortFive(ops)

	if !stack.IsSorted(ops.StackA) {
		t.Error("Stack should be sorted")
	}
	if ops.StackB != nil {
		t.Error("Stack B should be empty")
	}
}

func TestSortFiveRandom(t *testing.T) {
	tests := [][]int{
		{5, 4, 3, 2, 1},
		{3, 1, 4, 2, 5},
		{2, 5, 1, 4, 3},
	}

	for _, tt := range tests {
		ops := stack.NewOperations(true)
		for i := len(tt) - 1; i >= 0; i-- {
			ops.PushToA(tt[i])
		}

		SortFive(ops)

		if !stack.IsSorted(ops.StackA) {
			t.Errorf("Stack should be sorted for input %v", tt)
		}
		if ops.StackB != nil {
			t.Error("Stack B should be empty")
		}
	}
}

func TestSortSmallStack(t *testing.T) {
	ops := stack.NewOperations(true)
	ops.PushToA(1)
	ops.PushToA(2)

	Sort(ops)

	if !stack.IsSorted(ops.StackA) {
		t.Error("Stack should be sorted")
	}
}

func TestSortAlreadySorted(t *testing.T) {
	ops := stack.NewOperations(true)
	for i := 10; i >= 1; i-- {
		ops.PushToA(i)
	}

	Sort(ops)

	if !stack.IsSorted(ops.StackA) {
		t.Error("Stack should remain sorted")
	}
}

func TestSortLargeStack(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{"10 elements", []int{5, 2, 8, 1, 9, 3, 7, 4, 6, 10}},
		{"20 elements", []int{15, 3, 18, 7, 12, 1, 20, 9, 14, 5, 11, 8, 16, 2, 19, 6, 13, 4, 17, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ops := stack.NewOperations(true)
			for i := len(tt.values) - 1; i >= 0; i-- {
				ops.PushToA(tt.values[i])
			}

			Sort(ops)

			if !stack.IsSorted(ops.StackA) {
				t.Error("Stack should be sorted")
			}
			if ops.StackB != nil {
				t.Error("Stack B should be empty")
			}
		})
	}
}

func TestSortNegativeNumbers(t *testing.T) {
	ops := stack.NewOperations(true)
	values := []int{-5, 3, -1, 7, -10, 0, 2}
	for i := len(values) - 1; i >= 0; i-- {
		ops.PushToA(values[i])
	}

	Sort(ops)

	if !stack.IsSorted(ops.StackA) {
		t.Error("Stack should be sorted")
	}
	if ops.StackB != nil {
		t.Error("Stack B should be empty")
	}
}

func TestSort100Elements(t *testing.T) {
	ops := stack.NewOperations(true)

	values := []int{50, 25, 75, 12, 38, 62, 88, 6, 19, 31, 44, 56, 69, 81, 94,
		3, 9, 15, 22, 28, 35, 41, 47, 53, 59, 66, 72, 78, 85, 91, 97,
		1, 4, 7, 10, 13, 16, 20, 23, 26, 29, 32, 36, 39, 42, 45, 48,
		51, 54, 57, 60, 63, 67, 70, 73, 76, 79, 82, 86, 89, 92, 95, 98,
		2, 5, 8, 11, 14, 17, 21, 24, 27, 30, 33, 37, 40, 43, 46, 49,
		52, 55, 58, 61, 64, 68, 71, 74, 77, 80, 83, 87, 90, 93, 96, 99, 100,
		18, 34, 65, 84}

	for i := len(values) - 1; i >= 0; i-- {
		ops.PushToA(values[i])
	}

	Sort(ops)

	if !stack.IsSorted(ops.StackA) {
		t.Error("Stack with 100 elements should be sorted")
	}
	if ops.StackB != nil {
		t.Error("Stack B should be empty")
	}

	if stack.Size(ops.StackA) != 100 {
		t.Errorf("Expected stack size 100, got %d", stack.Size(ops.StackA))
	}
}
