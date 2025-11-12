package main

import (
	"fmt"
	"os"

	"github.com/kwa0x2/GoSortStack/internal/algorithm"
	"github.com/kwa0x2/GoSortStack/internal/parser"
	"github.com/kwa0x2/GoSortStack/internal/stack"
	"github.com/kwa0x2/GoSortStack/internal/visualizer"
)

func main() {
	args := os.Args[1:]

	visualize := false
	filteredArgs := []string{}
	for _, arg := range args {
		if arg == "-v" || arg == "--visualize" {
			visualize = true
		} else {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	ops, err := parser.Parse(filteredArgs, visualize)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if visualize {
		visualizer.Enable()
	}

	visualizer.ShowInitial(ops)

	if stack.IsSorted(ops.StackA) {
		if visualize {
			visualizer.ShowSummary()
		}
		return
	}

	size := stack.Size(ops.StackA)

	switch {
	case size <= 1:
		if visualize {
			visualizer.ShowSummary()
		}
		return
	case size == 2:
		if ops.StackA.Value > ops.StackA.Next.Value {
			if !visualize {
				fmt.Println("sa")
			}
			stack.SwapStack(&ops.StackA)
			visualizer.ShowStep(ops, "sa")
		}
	case size == 3:
		algorithm.SortThree(ops)
	case size <= 5:
		algorithm.SortFive(ops)
	default:
		algorithm.Sort(ops)
	}

	if visualize {
		visualizer.ShowSummary()
	} else {
		fmt.Println("\n--- Sorted Stack ---")
		ops.PrintA("Stack A")
		ops.PrintB("Stack B")
	}
}
