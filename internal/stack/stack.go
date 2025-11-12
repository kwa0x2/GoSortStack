package stack

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Operations struct {
	StackA *Node
	StackB *Node
	Silent bool
}

func NewOperations(silent bool) *Operations {
	return &Operations{
		StackA: nil,
		StackB: nil,
		Silent: silent,
	}
}

// add a value to the top of stack A
func (o *Operations) PushToA(value int) {
	newNode := &Node{Value: value, Next: o.StackA}
	o.StackA = newNode
}

// add a value to the top of stack B
func (o *Operations) PushToB(value int) {
	newNode := &Node{Value: value, Next: o.StackB}
	o.StackB = newNode
}

// debug
func (o *Operations) PrintA(label string) {
	fmt.Print(label, ": ")
	if o.StackA == nil {
		fmt.Println("(empty)")
		return
	}
	current := o.StackA
	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func (o *Operations) PrintB(label string) {
	fmt.Print(label, ": ")
	if o.StackB == nil {
		fmt.Println("(empty)")
		return
	}
	current := o.StackB
	for current != nil {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

func SwapStack(head **Node) {
	first := *head
	second := (*head).Next

	first.Next = second.Next
	second.Next = first
	*head = second
}

func RotateStack(head **Node) {
	first := *head
	last := GetLast(*head)

	*head = first.Next
	first.Next = nil
	last.Next = first
}

func ReverseRotateStack(head **Node) {
	i := 0
	tmp := *head

	for (*head).Next != nil {
		*head = (*head).Next
		i++
	}

	(*head).Next = tmp

	for i > 1 {
		tmp = tmp.Next
		i--
	}
	tmp.Next = nil
}

func PushStack(dst, src **Node) {
	tmp := *dst
	*dst = *src
	*src = (*src).Next
	(*dst).Next = tmp
}

func GetLast(head *Node) *Node {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	return current
}

func Size(head *Node) int {
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func Min(head *Node) int {
	if head == nil {
		return 0
	}

	min := head.Value
	current := head.Next
	for current != nil {
		if current.Value < min {
			min = current.Value
		}
		current = current.Next
	}
	return min
}

func Max(head *Node) int {
	if head == nil {
		return 0
	}

	max := head.Value
	current := head.Next
	for current != nil {
		if current.Value > max {
			max = current.Value
		}
		current = current.Next
	}
	return max
}

func FindIndex(head *Node, value int) int {
	index := 0
	current := head
	for current != nil {
		if current.Value == value {
			return index
		}
		index++
		current = current.Next
	}
	return -1
}
