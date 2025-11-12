package operations

import (
	"fmt"

	"github.com/kwa0x2/GoSortStack/internal/stack"
)

func SwapA(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil {
		return
	}
	stack.SwapStack(&o.StackA)
	if !o.Silent {
		fmt.Println("sa")
	}
}

func SwapB(o *stack.Operations) {
	if o.StackB == nil || o.StackB.Next == nil {
		return
	}
	stack.SwapStack(&o.StackB)
	if !o.Silent {
		fmt.Println("sb")
	}
}

func RotateA(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil {
		return
	}
	stack.RotateStack(&o.StackA)
	if !o.Silent {
		fmt.Println("ra")
	}
}

func RotateB(o *stack.Operations) {
	if o.StackB == nil || o.StackB.Next == nil {
		return
	}
	stack.RotateStack(&o.StackB)
	if !o.Silent {
		fmt.Println("rb")
	}
}

func PushA(o *stack.Operations) {
	if o.StackB == nil {
		return
	}
	stack.PushStack(&o.StackA, &o.StackB)
	if !o.Silent {
		fmt.Println("pa")
	}
}

func PushB(o *stack.Operations) {
	if o.StackA == nil {
		return
	}
	stack.PushStack(&o.StackB, &o.StackA)
	if !o.Silent {
		fmt.Println("pb")
	}
}

func ReverseRotateA(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil {
		return
	}
	stack.ReverseRotateStack(&o.StackA)
	if !o.Silent {
		fmt.Println("rra")
	}
}

func ReverseRotateB(o *stack.Operations) {
	if o.StackB == nil || o.StackB.Next == nil {
		return
	}
	stack.ReverseRotateStack(&o.StackB)
	if !o.Silent {
		fmt.Println("rrb")
	}
}

func ReverseRotateBoth(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil || o.StackB == nil || o.StackB.Next == nil {
		return
	}

	stack.ReverseRotateStack(&o.StackA)
	stack.ReverseRotateStack(&o.StackB)

	if !o.Silent {
		fmt.Println("rrr")
	}
}

func RotateBoth(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil || o.StackB == nil || o.StackB.Next == nil {
		return
	}

	stack.RotateStack(&o.StackA)
	stack.RotateStack(&o.StackB)

	if !o.Silent {
		fmt.Println("rr")
	}
}

func SwapBoth(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil || o.StackB == nil || o.StackB.Next == nil {
		return
	}

	stack.SwapStack(&o.StackA)
	stack.SwapStack(&o.StackB)

	if !o.Silent {
		fmt.Println("ss")
	}
}
 