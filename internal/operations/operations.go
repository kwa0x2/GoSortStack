package operations

import (
	"fmt"

	"github.com/kwa0x2/GoSortStack/internal/stack"
	"github.com/kwa0x2/GoSortStack/internal/visualizer"
)

func SwapA(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil {
		return
	}
	stack.SwapStack(&o.StackA)
	if !o.Silent {
		fmt.Println("SwapA")
	}
	visualizer.ShowStep(o, "SwapA")
}

func SwapB(o *stack.Operations) {
	if o.StackB == nil || o.StackB.Next == nil {
		return
	}
	stack.SwapStack(&o.StackB)
	if !o.Silent {
		fmt.Println("SwapB")
	}
	visualizer.ShowStep(o, "SwapB")
}

func RotateA(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil {
		return
	}
	stack.RotateStack(&o.StackA)
	if !o.Silent {
		fmt.Println("RotateA")
	}
	visualizer.ShowStep(o, "RotateA")
}

func RotateB(o *stack.Operations) {
	if o.StackB == nil || o.StackB.Next == nil {
		return
	}
	stack.RotateStack(&o.StackB)
	if !o.Silent {
		fmt.Println("RotateB")
	}
	visualizer.ShowStep(o, "RotateB")
}

func PushA(o *stack.Operations) {
	if o.StackB == nil {
		return
	}
	stack.PushStack(&o.StackA, &o.StackB)
	if !o.Silent {
		fmt.Println("PushA")
	}
	visualizer.ShowStep(o, "PushA")
}

func PushB(o *stack.Operations) {
	if o.StackA == nil {
		return
	}
	stack.PushStack(&o.StackB, &o.StackA)
	if !o.Silent {
		fmt.Println("PushB")
	}
	visualizer.ShowStep(o, "PushB")
}

func ReverseRotateA(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil {
		return
	}
	stack.ReverseRotateStack(&o.StackA)
	if !o.Silent {
		fmt.Println("ReverseRotateA")
	}
	visualizer.ShowStep(o, "ReverseRotateA")
}

func ReverseRotateB(o *stack.Operations) {
	if o.StackB == nil || o.StackB.Next == nil {
		return
	}
	stack.ReverseRotateStack(&o.StackB)
	if !o.Silent {
		fmt.Println("ReverseRotateB")
	}
	visualizer.ShowStep(o, "ReverseRotateB")
}

func ReverseRotateBoth(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil || o.StackB == nil || o.StackB.Next == nil {
		return
	}

	stack.ReverseRotateStack(&o.StackA)
	stack.ReverseRotateStack(&o.StackB)

	if !o.Silent {
		fmt.Println("ReverseRotateBoth")
	}
	visualizer.ShowStep(o, "ReverseRotateBoth")
}

func RotateBoth(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil || o.StackB == nil || o.StackB.Next == nil {
		return
	}

	stack.RotateStack(&o.StackA)
	stack.RotateStack(&o.StackB)

	if !o.Silent {
		fmt.Println("RotateBoth")
	}
	visualizer.ShowStep(o, "RotateBoth")
}

func SwapBoth(o *stack.Operations) {
	if o.StackA == nil || o.StackA.Next == nil || o.StackB == nil || o.StackB.Next == nil {
		return
	}

	stack.SwapStack(&o.StackA)
	stack.SwapStack(&o.StackB)

	if !o.Silent {
		fmt.Println("SwapBoth")
	}
	visualizer.ShowStep(o, "SwapBoth")
}
