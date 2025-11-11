package main

import (
	"fmt"
	"os"

	"github.com/kwa0x2/GoSortStack/internal/algorithm"
	"github.com/kwa0x2/GoSortStack/internal/checker"
	"github.com/kwa0x2/GoSortStack/internal/parser"
	"github.com/kwa0x2/GoSortStack/internal/stack"
)

func main() {
	args := os.Args[1:]

	ops, err := parser.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if checker.IsSorted(ops.StackA) {
		return
	}

	size := stack.Size(ops.StackA)

	switch {
	case size <= 1:
		return
	case size == 2:
		if ops.StackA.Value > ops.StackA.Next.Value {
			fmt.Println("sa")
		}
	case size == 3:
		algorithm.SortThree(ops)
	case size <= 5:
		algorithm.SortFive(ops)
	default:
		algorithm.Sort(ops)
	}

	fmt.Println("\n--- Sorted Stack ---")
	ops.PrintA("Stack A")
	ops.PrintB("Stack B")
}
