package parser

import (
	"testing"
)

func TestParseMultipleArgs(t *testing.T) {
	args := []string{"3", "2", "1"}
	ops, err := Parse(args, false)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ops.StackA.Value != 3 {
		t.Errorf("Expected top value 3, got %d", ops.StackA.Value)
	}
}

func TestParseQuotedString(t *testing.T) {
	args := []string{"3 2 1"}
	ops, err := Parse(args, false)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ops.StackA.Value != 3 {
		t.Errorf("Expected top value 3, got %d", ops.StackA.Value)
	}
}

func TestParseNegativeNumbers(t *testing.T) {
	args := []string{"-5", "10", "-3"}
	ops, err := Parse(args, false)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ops.StackA.Value != -5 {
		t.Errorf("Expected top value -5, got %d", ops.StackA.Value)
	}
}

func TestParseNoArgs(t *testing.T) {
	args := []string{}
	_, err := Parse(args, false)

	if err != ErrNoArgs {
		t.Errorf("Expected ErrNoArgs, got %v", err)
	}
}

func TestParseDuplicates(t *testing.T) {
	args := []string{"1", "2", "1"}
	_, err := Parse(args, false)

	if err != ErrDuplicate {
		t.Errorf("Expected ErrDuplicate, got %v", err)
	}
}

func TestParseInvalidInput(t *testing.T) {
	args := []string{"abc"}
	_, err := Parse(args, false)

	if err != ErrInvalidInput {
		t.Errorf("Expected ErrInvalidInput, got %v", err)
	}
}

func TestParseEmptyString(t *testing.T) {
	args := []string{""}
	_, err := Parse(args, false)

	if err != ErrNoArgs {
		t.Errorf("Expected ErrNoArgs, got %v", err)
	}
}

func TestParseSilentMode(t *testing.T) {
	args := []string{"1", "2", "3"}
	ops, err := Parse(args, true)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !ops.Silent {
		t.Error("Expected Silent to be true")
	}
}

func TestParseWithPlusSign(t *testing.T) {
	args := []string{"+5", "+10"}
	ops, err := Parse(args, false)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if ops.StackA.Value != 5 {
		t.Errorf("Expected top value 5, got %d", ops.StackA.Value)
	}
}

func TestParseInvalidNumber(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{"Only plus sign", []string{"+"}},
		{"Only minus sign", []string{"-"}},
		{"Letters", []string{"12a"}},
		{"Special chars", []string{"12@34"}},
		{"Decimal", []string{"12.34"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.input, false)
			if err == nil {
				t.Error("Expected error for invalid input")
			}
		})
	}
}
