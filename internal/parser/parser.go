package parser

import (
	"errors"
	"strconv"
	"strings"

	"github.com/kwa0x2/GoSortStack/internal/stack"
)

var (
	ErrNoArgs       = errors.New("no args")
	ErrInvalidInput = errors.New("invalid input")
	ErrDuplicate    = errors.New("duplicate")
	ErrNotInt       = errors.New("not int")
)


func Parse(args []string) (*stack.Operations, error) {
	if len(args) == 0 {
		return nil, ErrNoArgs
	}

	var numbers []int

	if len(args) == 1 {
		nums, err := parseQuotedString(args[0])
		if err != nil {
			return nil, err
		}
		numbers = nums
	} else {
		nums, err := parseMultipleArgs(args)
		if err != nil {
			return nil, err
		}
		numbers = nums
	}

	if hasDuplicates(numbers) {
		return nil, ErrDuplicate
	}

	ops := stack.NewOperations(false)
	loadStack(ops, numbers)

	return ops, nil
}

func parseQuotedString(s string) ([]int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, ErrNoArgs
	}

	parts := strings.Fields(s)
	return stringsToInts(parts)
}

func parseMultipleArgs(args []string) ([]int, error) {
	return stringsToInts(args)
}

func stringsToInts(strs []string) ([]int, error) {
	numbers := make([]int, 0, len(strs))

	for _, s := range strs {
		if !isValidNumber(s) {
			return nil, ErrInvalidInput
		}

		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, ErrNotInt
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}


func isValidNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0
	if s[i] == '+' || s[i] == '-' {
		i++
		if i >= len(s) {
			return false
		}
	}

	for i < len(s) {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
		i++
	}

	return true
}

func hasDuplicates(numbers []int) bool {
	seen := make(map[int]bool)
	for _, num := range numbers {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}


func loadStack(ops *stack.Operations, numbers []int) {
	for i := len(numbers) - 1; i >= 0; i-- {
		ops.PushToA(numbers[i])
	}
}
